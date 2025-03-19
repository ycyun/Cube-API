package config

import (
	"github.com/goccy/go-json"
	"os"
	"time"
)

type StructConfig struct {
	ServerPort   string        `json:"server_port"`
	TaskInterval time.Duration `json:"task_interval"`
}

func Load() *StructConfig {
	ret := new(StructConfig)
	ret.ServerPort = "34585"

	return ret
}

const configFile = "./config.json"

func SaveConfig() {
	config := StructConfig{
		//Neighbor: controller.Neighbor.Neighbors,
	}

	fc, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0666)
	defer fc.Close()

	strconfig, err := json.Marshal(config)

	_, err = fc.Write(strconfig)
	if err != nil {
		return
	}

}
