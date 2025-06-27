package Base

import (
	"fmt"
	"log/slog"
	"reflect"
	"sync"
)

type Struct struct {
	ID string `json:"id"`
} //@name Struct

// PStruct BaseStruct 싱글톤 인스턴스로, 지연 초기화되며 스레드 안전성을 위해 sync.Once로 보호됩니다.
var lockSample sync.Once
var PStruct *Struct

func Init() *Struct {
	if PStruct == nil {
		lockSample.Do(
			func() {
				fmt.Println("Creating ", reflect.TypeOf(PStruct), " now.")
				PStruct = &Struct{}
			})
	} else {
		//fmt.Println("Creating ", reflect.TypeOf(PStruct), "with config", conf, " now.")
		slog.Debug("Use Old Struct", "type", reflect.TypeOf(PStruct))
	}
	return PStruct
}
