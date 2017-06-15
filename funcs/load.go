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

	L = append(L, CounterValue("load.load1", loadVal.Load1))
	L = append(L, CounterValue("load.load5", loadVal.Load5))
	L = append(L, CounterValue("load.load15", loadVal.Load15))

	return
}
