package server

import (
	"whatsapp-clone/db"
	"whatsapp-clone/es"
	config "whatsapp-clone/pkg/config/v1"
	ping "whatsapp-clone/pkg/ping/v1"
	sampleV1 "whatsapp-clone/pkg/sample/v1"
	sampleV2 "whatsapp-clone/pkg/sample/v2"

	"whatsapp-clone/redis"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func NewRouter() *gin.Engine {
	dbFactory := db.Factory
	redis := redis.Client()
	es := es.Client()

	router := gin.New()
	router.Use(apmgin.Middleware(router))
	router.Use(gin.Logger())

	pingGroup := router.Group("ping")
	pingGroup.GET("", ping.NewModule().GetController().Ping)

	v1 := router.Group("api/v1")
	{
		sampleGroupV1 := v1.Group("sample")
		{
			sampleControllerV1 := sampleV1.NewSampleModule(dbFactory, redis, es).GetController()
			sampleGroupV1.GET("/hello", sampleControllerV1.Hello)
			sampleGroupV1.GET("/error", sampleControllerV1.Error)
			sampleGroupV1.GET("/db", sampleControllerV1.FromDb)
			sampleGroupV1.GET("/redis", sampleControllerV1.FromRedis)
			sampleGroupV1.GET("/panic", sampleControllerV1.Panic)
			sampleGroupV1.GET("/es", sampleControllerV1.FromEs)
		}

		configGroup := v1.Group("config")
		{
			configGroupController := config.NewModuleSingleton().GetController()
			configGroup.POST("/apmsamplerate", configGroupController.APMSampleRate)
		}
	}

	v2 := router.Group("api/v2")
	{
		sampleGroupV2 := v2.Group("sample")
		{
			sampleV2Controller := sampleV2.NewSampleModule().GetController()
			sampleGroupV2.GET("/hello", sampleV2Controller.Hello)
		}
	}

	return router
}
