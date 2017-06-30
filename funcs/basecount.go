package funcs

import (
	"fmt"
	"log"

	"github.com/open-falcon/common/model"
	"github.com/oslet/agent/tools/wmi"
)

type Win32_PerfRawData_PerfProc_Process struct {
	Name        string
	HandleCount uint32
}

type Win32_PerfFormattedData struct {
	Processes uint64
	Threads   uint64
}

func ProcessCounters() ([]Win32_PerfFormattedData, error) {
	ret := make([]Win32_PerfFormattedData, 0)
	var dst []Win32_PerfFormattedData
	err := wmi.Query("SELECT Processes,Threads FROM Win32_PerfFormattedData_PerfOS_System", &dst)
	if err != nil {
		return ret, err
	}

	for _, d := range dst {

		ret = append(ret, Win32_PerfFormattedData{
			Processes: uint64(d.Processes),
			Threads:   uint64(d.Threads),
		})
	}

	return ret, nil
}

func HandleCounters() ([]Win32_PerfRawData_PerfProc_Process, error) {
	ret := make([]Win32_PerfRawData_PerfProc_Process, 0)
	var dst []Win32_PerfRawData_PerfProc_Process
	err := wmi.Query("select HandleCount from Win32_PerfFormattedData_PerfProc_Process", &dst)

	if err != nil {
		return ret, err
	}

	for _, d := range dst {

		ret = append(ret, Win32_PerfRawData_PerfProc_Process{
			Name:        string(d.Name),
			HandleCount: uint32(d.HandleCount),
		})
	}
	return ret, nil
}

func ProcessMetrics() (L []*model.MetricValue) {

	dsList, err := ProcessCounters()
	if err != nil {
		log.Println("Get process data fail: ", err)
		return
	}

	for _, ds := range dsList {
		L = append(L, GaugeValue("process.total", ds.Processes))
		L = append(L, GaugeValue("thread.total", ds.Threads))
	}
	return
}

func HandleCountMetrics() (L []*model.MetricValue) {

	Handle, err := HandleCounters()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, hc := range Handle {
		if hc.Name == "_Total" {
			L = append(L, GaugeValue("handle.total", hc.HandleCount))
		}
	}
	return
}
