package main

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/ycyun/Cube-API/controller"
	"github.com/ycyun/Cube-API/utils"
)

//	@title			Cube APIServers
//	@version		1.0
//	@description	This is a Cube-APIServer server.
//	@termsOfService	https://ablecloud.io/

//	@contact.name	APIServer Support
//	@contact.url	https://www.ablecloud.io/support
//	@contact.email	ycyun@ablecloud.io

//	@license.name	Apache 2.0
//	@license.url	https://www.apache.org/licenses/LICENSE-2.0.html

//	@ssshost						10.211.55.11:8080
//	@BasePath					/api/v1
//	@Schemes					http https
//	@securityDefinitions.basic	None

// @externalDocs.description	ABLECLOUD
// @externalDocs.url			https://www.ablecloud.io
func main() {
	// 시간대 설정
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic(err)
	}
	// Set the timezone for the current process
	time.Local = location
	//Logger := utils.LogInit()
	slog.SetDefault(utils.Logger) // default 설정. logger 대신 slog로 로그 찍어도 logger랑 똑같은 기능을 함.
	slog.SetLogLoggerLevel(slog.LevelError)
	Logger := utils.Logger
	Logger.Debug("server on ..")
	Logger.Debug("controller init")

	Controller := controller.Init()

	API := Controller.A
	Worker := Controller.W
	Logger.Info("handlers", "handler info", Worker.Handlers[0])
	go Worker.Run()
	API.Run()
	//
	////Controller.StatusRegister(Mold.MonitorStatus)
	//Controller.StatusRegister(Glue.Monitor)
	////Controller.StatusRegister(Dashboard.Monitor)
	//Controller.StatusRegister(PCS.Monitor)
	//Controller.StatusRegister(Cube.Hosts.Update)
	//Controller.StatusRegister(Cube.NICs.Update)
	//Controller.StatusRegister(Cube.Disks.Update)
	//Controller.StatusRegister(C.SaveConfig)
	//
	//go Controller.Start()
	//APIPort := "8080"
	//docs.SwaggerInfo.Schemes = []string{"http", "https"}
	//docs.SwaggerInfo.Host = UTILS.GetLocalIP().String() + ":" + APIPort
	//log.SetFlags(log.LstdFlags | log.Lshortfile)
	//
	//Controller.Stop()
	fmt.Println("end")
}
