package Sample

import (
	"fmt"
	"reflect"
	"sync"
)

type StructSamples struct {
	Value []string `json:"value"`
} //	@name	StructSamples

var lockSample sync.Once
var samples *StructSamples

func Init() *StructSamples {
	if samples == nil {
		lockSample.Do(
			func() {
				fmt.Println("Creating ", reflect.TypeOf(samples), " now.")
				samples = &StructSamples{}
			})
	} else {
		fmt.Println("get old ", reflect.TypeOf(samples), " instance.")
	}
	return samples
}
