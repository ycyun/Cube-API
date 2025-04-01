package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ycyun/Cube-API/Modules/api"    // APIRouter 핸들러 임포트
	"github.com/ycyun/Cube-API/Modules/config" // 설정 임포트
	"log"
	"reflect"
	"sync"
)

type APIRouter struct {
	Config *config.StructConfig `json:"cfg"`
}

var lockRouter sync.Once //Router 생성에 대한 lock
var PRouter *APIRouter   //Router의 포인터

func Init(conf *config.StructConfig) *APIRouter {
	if PRouter == nil {
		lockRouter.Do(
			func() {
				fmt.Println("Creating ", reflect.TypeOf(PRouter), "with config", conf, " now.")
				PRouter = &APIRouter{
					Config: conf,
				}
			})
	} else {
		fmt.Println("get old ", reflect.TypeOf(PRouter), " instance.")
	}
	return PRouter
}

func (*APIRouter) Run() {
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

	// APIRouter 엔드포인트 등록
	api.RegisterRoutes(router)

	// 서버 실행
	log.Println("Starting APIRouter server on port", cfg.ServerPort)
	if err = router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
		//c.AddError(err)
	}
}
