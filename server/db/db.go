package db

import (
	"database/sql"
	"sync"
	"time"

	"whatsapp-clone/config"

	"whatsapp-clone/logger"

	"go.elastic.co/apm/module/apmsql"

	_ "go.elastic.co/apm/module/apmsql/mysql"
	"go.uber.org/zap"
)

var reader *sql.DB
var writer *sql.DB
var once sync.Once

type DBConfig struct {
	DBUserName           string
	DBPassword           string
	DBHost               string
	DBPort               string
	DBName               string
	DBMaxIdleConnections int
	DBMaxOpenConnections int
	DBConnMaxLifetime    time.Duration
}

func NewDBClient(config *DBConfig) *sql.DB {
	url := config.DBUserName + ":" + config.DBPassword + "@tcp(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?multiStatements=true&parseTime=true"
	client, err := apmsql.Open("mysql", url)
	if err != nil {
		panic(err.Error())
	}

	client.SetMaxIdleConns(config.DBMaxIdleConnections)
	client.SetMaxOpenConns(config.DBMaxOpenConnections)
	client.SetConnMaxLifetime(time.Minute * 10)
	return client
}

func Init() {
	once.Do(func() {
		config := config.Get()

		writerConfig := &DBConfig{
			DBUserName:           config.DBUserName,
			DBPassword:           config.DBPassword,
			DBHost:               config.DBHostWriter,
			DBPort:               config.DBPort,
			DBName:               config.DBName,
			DBMaxIdleConnections: config.DBMaxIdleConnections,
			DBMaxOpenConnections: config.DBMaxOpenConnections,
			DBConnMaxLifetime:    time.Minute * 10,
		}

		readerConfig := writerConfig
		readerConfig.DBHost = config.DBHostReader

		reader = NewDBClient(readerConfig)
		writer = NewDBClient(writerConfig)

		logger.Client().Info("writer connected", zap.String("host", config.DBHostReader))
		logger.Client().Info("reader connected", zap.String("host", config.DBHostWriter))
	})
}

func Factory(typ string) *sql.DB {
	switch typ {
	case "reader":
		return reader
	case "writer":
		return writer
	default:
		panic("no such db")
	}
}

func WrapQuery(query string) string {
	return config.Get().SqlPrefix + query
}

type DBFactory func(t string) *sql.DB