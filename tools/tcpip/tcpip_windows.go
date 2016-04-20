package tcpip

import (

	//	"bytes"
	//	"syscall"
	//	"unsafe"

	"github.com/LeonZYang/agent/tools/wmi"

	//	"github.com/LeonZYang/agent/tools/internal/common"
)

type Win32_PerfFormattedData struct {
	ConnectionFailures     uint64
	ConnectionsActive      uint64
	ConnectionsPassive     uint64
	ConnectionsEstablished uint64
	ConnectionsReset       uint64
}

func TcpipCounters() ([]Tcpipdatastat, error) {
	ret := make([]Tcpipdatastat, 0)
	var dst []Win32_PerfFormattedData
	err := wmi.Query("SELECT ConnectionFailures,ConnectionsActive,ConnectionsPassive,ConnectionsEstablished,ConnectionsReset FROM Win32_PerfRawData_Tcpip_TCPv4", &dst)
	if err != nil {
		return ret, err
	}

	for _, d := range dst {

		ret = append(ret, Tcpipdatastat{
			ConFailures:    uint64(d.ConnectionFailures),
			ConActive:      uint64(d.ConnectionsActive),
			ConPassive:     uint64(d.ConnectionsPassive),
			ConEstablished: uint64(d.ConnectionsEstablished),
			ConReset:       uint64(d.ConnectionsReset),
		})
	}

	return ret, nil
}
