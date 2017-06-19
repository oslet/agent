package funcs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/open-falcon/common/model"
)

const (
	winPre   = "cmd /c "
	LinuxPre = "bash /c "
)

type Socketstat struct {
	Established uint64 `json:"established"`
	Close_wait  uint64 `json:"close_wait"`
	Time_wait   uint64 `json:"time_wait"`
}

type SocketCount struct {
	ConEstablished uint64
	ConClose_wait  uint64
	ConTime_wait   uint64
}

func (d Socketstat) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}

func SocketMetrics() []*model.MetricValue {

	cmd := `netstat -n`
	result, err := RunCmd(cmd)
	if err != nil {
		log.Fatal(err)
	}

	a := string(result)

	var es, cw, tw int
	es = strings.Count(a, "ESTABLISHED")
	cw = strings.Count(a, "CLOSE_WAIT")
	tw = strings.Count(a, "TIME_WAIT")

	return []*model.MetricValue{
		GaugeValue("ss.estab", uint64(es)),
		GaugeValue("ss.closed", uint64(cw)),
		GaugeValue("ss.timewait", uint64(tw)),
	}

}

func RunCmd(cmd string) (result []byte, err error) {
	cmdPip := strings.Split(cmd, "|")
	if len(cmdPip) < 2 {
		return run(cmd).Output()
	}

	return runPipe(cmdPip)
}

func run(cmd string) *exec.Cmd {
	if strings.Contains(cmd, "findstr") || !strings.Contains(cmd, "find") {
		cmd = strings.Replace(cmd, `"`, "", -1)
		cmd = strings.TrimSpace(cmd)
	}
	cmdList := strings.Split(cmd, " ")

	return exec.Command(cmdList[0], cmdList[1:]...)
}

// 不支持windows的find命令，原因未知
func runPipe(pip []string) (result []byte, err error) {
	var cmds []*exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmds = append(cmds, run(winPre+pip[0]))
	case "linux":
		cmds = append(cmds, run(LinuxPre+pip[0]))
	default:
		return nil, fmt.Errorf("Not supported by the system. %v", runtime.GOOS)
	}

	for i := 1; i < len(pip); i++ {
		cmds = append(cmds, run(pip[i]))
		cmds[i].Stdin, err = cmds[i-1].StdoutPipe()
		if err != nil {
			return nil, err
		}
	}

	end := len(cmds) - 1
	// cmds[end].Stdout = os.Stdout
	stdout, _ := cmds[end].StdoutPipe()

	for i := end; i > 0; i-- {
		cmds[i].Start()
	}
	cmds[0].Run()

	buf := make([]byte, 102400000)
	n, _ := stdout.Read(buf)

	err = cmds[end].Wait()

	return buf[:n], err
}

func readPipe() {
	var err error
	cmd1 := exec.Command("cmd", "/c", `echo hello world.`)
	cmd2 := exec.Command("find", "/i", `"he"`)
	cmd2.Stdin, err = cmd1.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	cmd2.Start()
	cmd1.Run()
	err = cmd2.Wait()
	fmt.Println(err)
}
