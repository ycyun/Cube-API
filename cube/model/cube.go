package model

import (
	"fmt"
	"github.com/ycyun/Cube-API/utils"
	"reflect"
	"sync"
	"time"
)

type TypeCUBE struct {
	Handlers    []func()          `json:"handlers"`
	running     bool              `json:"running"`
	Neighbor    []TypeHost        `json:"neighbors"`
	errors      []utils.Errorlog  `json:"errors"`
	version     utils.TypeVersion `json:"version"`
	RefreshTime time.Time         `json:"refreshTime"`
}

var lockCUBE sync.Once
var cube *TypeCUBE

func Cube() *TypeCUBE {
	if cube == nil {
		lockCUBE.Do(
			func() {
				fmt.Println("Creating ", reflect.TypeOf(cube), " now.")
				cube = &TypeCUBE{
					Handlers: []func(){},
					running:  false,
					version:  utils.TypeVersion{Version: "v5.0.0"},
				}
			})
	} else {
		fmt.Println("get old ", reflect.TypeOf(cube), " instance.")
	}

	return cube
}

func (c *TypeCUBE) GetVersion() utils.TypeVersion {
	return c.version
} // @name version
