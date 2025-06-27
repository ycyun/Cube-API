package utils

import (
	"log/slog"
	"reflect"
	"sync"
	"time"
)

var lockConfig sync.Once  // Worker의 락
var Pconfig *StructConfig // Worker의 포인터

const configFile = "./config.json"

type StructConfig struct {
	ServerPort   string        `json:"server_port"`
	TaskInterval time.Duration `json:"task_interval"`
	LogDir       string        `json:"log_dir"`
}

func Init(conf *StructConfig) *StructConfig {
	if Pconfig == nil {
		lockConfig.Do(
			func() {
				//fmt.Println("Creating ", reflect.TypeOf(Pconfig), "with config", conf, " now.")
				slog.Debug("Create Struct", "type", reflect.TypeOf(Pconfig), "config", conf)
				Pconfig = &StructConfig{
					ServerPort:   conf.ServerPort,
					TaskInterval: conf.TaskInterval,
					LogDir:       conf.LogDir,
				}
			})
	} else {
		//fmt.Println("Creating ", reflect.TypeOf(Pconfig), "with config", conf, " now.")
		slog.Debug("Use Old Struct", "type", reflect.TypeOf(Pconfig), "config", conf)
	}
	return Pconfig
}
