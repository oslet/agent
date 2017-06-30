package funcs

import (
	"log"

	"github.com/open-falcon/common/model"
	"github.com/oslet/agent/tools/load"
)

func LoadMetrics() (L []*model.MetricValue) {

	loadVal, err := load.LoadAvg()
	if err != nil {
		log.Println("Get load fail: ", err)
		return nil
	}

	L = append(L, CounterValue("load.load1min", loadVal.Load1min))
	L = append(L, CounterValue("load.load5min", loadVal.Load5min))
	L = append(L, CounterValue("load.load15min", loadVal.Load15min))

	return
}
