package funcs

import (
	"github.com/open-falcon/common/model"
	"github.com/oslet/agent/g"
)

type FuncsAndInterval struct {
	Fs       []func() []*model.MetricValue
	Interval int
}

var Mappers []FuncsAndInterval

func BuildMappers() {
	interval := g.Config().Transfer.Interval
	Mappers = []FuncsAndInterval{
		FuncsAndInterval{
			Fs: []func() []*model.MetricValue{
				AgentMetrics,
				NetMetrics,
				CpuMetrics,
				DeviceMetrics,
				DiskIOMetrics,
				LoadMetrics,
				MemMetrics,
				TcpipMetrics,
			},
			Interval: interval,
		},
	}
}
