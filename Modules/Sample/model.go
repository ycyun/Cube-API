package Sample

import (
	"fmt"
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
				fmt.Println("Creating ", reflect.TypeOf(Psamples), " now.")
				Psamples = &StructSamples{}
			})
	} else {
		fmt.Println("get old ", reflect.TypeOf(Psamples), " instance.")
	}
	return Psamples
}
