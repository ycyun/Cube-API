package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ycyun/Cube-API/Modules/api"    // APIServer 핸들러 임포트
	"github.com/ycyun/Cube-API/Modules/config" // 설정 임포트
	"log"
	"reflect"
	"sync"
)

type APIServer struct {
	Config *config.StructConfig `json:"cfg"`
}

var lockAPI sync.Once
var Papi *APIServer

func Init(conf *config.StructConfig) *APIServer {
	if Papi == nil {
		lockAPI.Do(
			func() {
				fmt.Println("Creating ", reflect.TypeOf(Papi), "with config", conf, " now.")
				Papi = &APIServer{
					Config: conf,
				}
			})
	} else {
		fmt.Println("get old ", reflect.TypeOf(Papi), " instance.")
	}
	return Papi
}

func (*APIServer) Run() {
	// 환경 변수 및 설정 로드
	var err error
	cfg := config.Load()

	// Gin 엔진 생성
	router := gin.Default()
	router.Use(gin.Logger())
	// Recovery 미들웨어는 panic이 발생하면 500 에러를 씁니다.
	router.Use(gin.Recovery())

	//gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)
	router.ForwardedByClientIP = true
	err = router.SetTrustedProxies(nil)

	// APIServer 엔드포인트 등록
	api.RegisterRoutes(router)

	// 서버 실행
	log.Println("Starting APIServer server on port", cfg.ServerPort)
	if err = router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
		//c.AddError(err)
	}
}
