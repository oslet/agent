package funcs

import (
	"fmt"

	"github.com/oslet/agent/tools/cpu"
)

func CheckCollector() {

	output := make(map[string]bool)

	_, procStatErr := cpu.CPUTimes(false)

	output["df.bytes"] = len(DeviceMetrics()) > 0
	output["net.if  "] = len(CoreNetMetrics()) > 0
	output["loadavg "] = len(LoadMetrics()) > 0
	output["cpustat "] = procStatErr == nil
	output["disk.io "] = len(DiskIOMetrics()) > 0
	output["memory  "] = len(MemMetrics()) > 0
	output["tcpip   "] = len(TcpipMetrics()) > 0
	output["proc    "] = len(ProcessMetrics()) > 0
	output["handle  "] = len(HandleCountMetrics()) > 0

	for k, v := range output {
		status := "fail"
		if v {
			status = "ok"
		}
		fmt.Println(k, "...", status)
	}
}
