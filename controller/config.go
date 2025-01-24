package controller

import (
	"github.com/goccy/go-json"
	"os"
)

type Config struct {
	Neighbor []TypeNeighbor `json:"neighbor"`
}

const configFile = "./config.json"

func SaveConfig() {
	config := Config{
		Neighbor: controller.Neighbor.Neighbors,
	}

	fc, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0666)
	defer fc.Close()

	strconfig, err := json.Marshal(config)

	_, err = fc.Write(strconfig)
	if err != nil {
		return
	}

}
