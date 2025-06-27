package Network

import (
	"log/slog"
	"reflect"
	"sync"
)

// Psamples는 StructSamples의 싱글톤 인스턴스로, 지연 초기화되며 스레드 안전성을 위해 sync.Once로 보호됩니다.
var lockNetwork sync.Once
var PNetwork *StructNetwork

type StructNetwork struct {
	Value []string `json:"value"`
} //	@name	StructNetwork

func Init() *StructNetwork {
	if PNetwork == nil {
		lockNetwork.Do(
			func() {

				slog.Debug("Create Struct", "type", reflect.TypeOf(PNetwork))
				PNetwork = &StructNetwork{}
			})
	} else {
		slog.Debug("Use Old Struct", "type", reflect.TypeOf(PNetwork))
	}
	return PNetwork
}
