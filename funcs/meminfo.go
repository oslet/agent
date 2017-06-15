package funcs

import (
	"log"

	"github.com/open-falcon/common/model"
	"github.com/oslet/agent/tools/mem"
)

func MemMetrics() []*model.MetricValue {
	m, err := mem.VirtualMemory()
	if err != nil {
		log.Println("Get memory fail: ", err)
		return nil
	}

	return []*model.MetricValue{
		GaugeValue("mem.memtotal", m.Total),
		GaugeValue("mem.memused", m.Used),
		GaugeValue("mem.memfree", m.Available),
		GaugeValue("mem.memfree.percent", 100.0-m.UsedPercent),
		GaugeValue("mem.memused.percent", m.UsedPercent),
	}

}
