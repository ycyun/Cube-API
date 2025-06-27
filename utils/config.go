package utils

import (
	"fmt"
	"github.com/goccy/go-json"
	"os"
	"strconv"
	"time"
)

func LoadConfig() *StructConfig {

	fc, err := os.OpenFile(configFile, os.O_RDONLY, 0666)
	defer fc.Close()
	if err != nil {
		ret := new(StructConfig)
		ret.ServerPort = "34585"
		ret.TaskInterval = time.Duration(60) * time.Second
		ret.LogDir = "./log"
		Init(ret)
	} else {
		stat, _ := fc.Stat()
		buff := make([]byte, stat.Size()+1)
		ret := new(StructConfig)
		n, errRead := fc.Read(buff)
		if errRead != nil {
			fmt.Println("read ", buff, " with ", strconv.Itoa(n), "bytes, err_read: ", errRead)
		}
		errUnmarshal := json.Unmarshal(buff, ret)
		if errUnmarshal != nil {
			fmt.Println("unmarshal ", buff, " with errUnmarshal: ", errUnmarshal)
		}
		Init(ret)
	}

	return Pconfig
}
func (Pconfig *StructConfig) GetServerPort() string {
	return Pconfig.ServerPort
}
func (Pconfig *StructConfig) GetTaskInterval() time.Duration {
	return Pconfig.TaskInterval
}
func (Pconfig *StructConfig) SetServerPort(port string) {
	Pconfig.ServerPort = port
	SaveConfig()
}
func (Pconfig *StructConfig) SetTaskInterval(interval time.Duration) {
	Pconfig.TaskInterval = interval
	SaveConfig()
}
func SaveConfig() {

	fc, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0666)
	defer fc.Close()

	strconfig, err := json.Marshal(Pconfig)

	_, err = fc.Write(strconfig)
	if err != nil {
		return
	}

}
