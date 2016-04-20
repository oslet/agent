package funcs

import (
	"log"

	"github.com/LeonZYang/agent/tools/tcpip"
	"github.com/open-falcon/common/model"
)

func TcpipMetrics() (L []*model.MetricValue) {

	dsList, err := tcpip.TcpipCounters()
	if err != nil {
		log.Println("Get tcpip data fail: ", err)
		return
	}

	for _, ds := range dsList {
		L = append(L, GaugeValue("tcpip.confailures", ds.ConFailures))
		L = append(L, GaugeValue("tcpip.conactive", ds.ConActive))
		L = append(L, GaugeValue("tcpip.conpassive", ds.ConPassive))
		L = append(L, GaugeValue("tcpip.conestablished", ds.ConEstablished))
		L = append(L, GaugeValue("tcpip.conreset", ds.ConReset))
	}
	return
}
