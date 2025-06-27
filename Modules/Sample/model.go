package Sample

import (
	"log/slog"
	"reflect"
	"sync"
)

type StructSamples struct {
	Value []string `json:"value"`
} //	@name	StructSamples

// Psamples는 StructSamples의 싱글톤 인스턴스로, 지연 초기화되며 스레드 안전성을 위해 sync.Once로 보호됩니다.
var lockSample sync.Once
var Psamples *StructSamples

func Init() *StructSamples {
	if Psamples == nil {
		lockSample.Do(
			func() {
				slog.Debug("Create Struct", "type", reflect.TypeOf(Psamples))
				Psamples = &StructSamples{}
			})
	} else {
		slog.Debug("Use Old Struct", "type", reflect.TypeOf(Psamples))
	}
	return Psamples
}
