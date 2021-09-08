package server

import (
	"trell/go-starter/db"
	// "trell/go-starter/es"
	"trell/go-starter/logger"
	"trell/go-starter/redis"
)

func Init() {
	logger.Init()
	db.Init()
	// es.Init()
	redis.Init()

	r := NewRouter()
	r.Run(":" + "4000")
}
