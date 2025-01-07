package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	Cube "github.com/ycyun/Cube-API/cube/action"
	Dashboard "github.com/ycyun/Cube-API/dashboard/action"
	"github.com/ycyun/Cube-API/docs"
	Glue "github.com/ycyun/Cube-API/glue/action"
	Mold "github.com/ycyun/Cube-API/mold/action"
	"log"
	"time"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
func main() {
	// 시간대 설정
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic(err)
	}
	// Set the timezone for the current process
	time.Local = location

	//cube := Cube.Init()
	Cube.StatusRegister(Mold.MonitorStatus)
	Cube.StatusRegister(Glue.Monitor)
	//Cube.StatusRegister(Glue.MonitorGlueStatus)
	//Cube.StatusRegister(Glue.MonitorGlueHealthDetail)
	Cube.StatusRegister(Dashboard.Monitor)
	//cube.StatusRegister(Dashboard.MonitorDashboard)

	Cube.StatusRegister(errorMaker)

	go Cube.Start()

	docs.SwaggerInfo.Title = "Cube API"
	docs.SwaggerInfo.Description = "This is a Cube-API server."
	docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = ".swagger.io"

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)
	r.ForwardedByClientIP = true
	err = r.SetTrustedProxies(nil)
	if err != nil {
		Cube.AddError(err)
	}

	r.Use(gin.Logger())

	// Recovery 미들웨어는 panic이 발생하면 500 에러를 씁니다.
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		glue := v1.Group("/glue")
		{
			glue.GET("", Glue.GetGlueStatus)
			glue.GET("/auth", Glue.GetGlueAuth)
			glue.GET("/auth/:username", Glue.GetGlueAuth)
			glue.GET("/auths", Glue.GetGlueAuths)
		}
		mold := v1.Group("/mold")
		{
			mold.GET("", Mold.GetStatus)
		}
		dashboard := v1.Group("/dashboard")
		{
			dashboard.GET("", Dashboard.UpdateStatus)

		}
		v1.Any("/version", Cube.Version)
		v1.GET("/err", Cube.Error)
		v1.DELETE("/err", Cube.ClearError)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err = r.Run(":8080")
	if err != nil {
		Cube.AddError(err)
	}
	Cube.Stop()
	Cube.PrintError()
	fmt.Println("end")
}

func errorMaker() {
	Cube.AddError(errors.New(time.Now().String()))
}
