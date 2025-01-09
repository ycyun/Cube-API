package controller

import (
	"bufio"
	"fmt"
	"github.com/goccy/go-json"
	"io"
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
func LoadConfig() {
	fc, err := os.OpenFile(configFile, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fc.Close()

	var strconfig []byte
	var config Config

	strconfig, err = io.ReadAll(bufio.NewReader(fc))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(strconfig, &config)

	if err != nil {
		fmt.Println(err)
		return
	}
	Init()
	controller.Neighbor.Neighbors = config.Neighbor
}
