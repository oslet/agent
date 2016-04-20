package tcpip

import (
	"encoding/json"
)

type Tcpipdatastat struct {
	ConFailures    uint64 `json:"confailures"`
	ConActive      uint64 `json:"conactive"`
	ConPassive     uint64 `json:"conpassive"`
	ConEstablished uint64 `json:"conestablished"`
	ConReset       uint64 `json:"conreset"`
}

func (d Tcpipdatastat) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}
