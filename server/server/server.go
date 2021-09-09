package server

import (
	"whatsapp-clone/db"
	// "whatsapp-clone/es"
	"whatsapp-clone/logger"
	"whatsapp-clone/redis"
)

func Init() {
	logger.Init()
	db.Init()
	// es.Init()
	redis.Init()

	r := NewRouter()
	r.Run(":" + "4000")
}
