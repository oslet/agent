package load

import (
	"encoding/json"
)

type LoadAvgStat struct {
	Load1min  float64 `json:"load1min"`
	Load5min  float64 `json:"load5min"`
	Load15min float64 `json:"load15min"`
}

func (l LoadAvgStat) String() string {
	s, _ := json.Marshal(l)
	return string(s)
}
