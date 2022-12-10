package main

import (
	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/nsrvel/go-fiber-boilerplate/internal/server"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/infra/db"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/infra/logger"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/infra/redis"
)

func main() {

	//* ====================== Config ======================

	conf := config.InitConfig("local")

	//* ====================== Logger ======================

	//* Loggrus
	appLogger := logger.NewLogrusLogger(&conf.Logger.Logrus)

	//* Grafana Loki
	if conf.Grafana.IsActive {
		if conf.App.Env != "local" {
			err := logger.InitLoki(conf, appLogger)
			if err != nil {
				appLogger.Errorf("Grafana Loki err: %s", err.Error())
			}
		}
	}

	//* ====================== Connection DB ======================

	var dbList db.DatabaseList

	//? SimpleTransaction DB
	dbList.SimpleTransaction = db.NewGORMConnection(&conf.Connection.SimpleTransaction, appLogger)

	//? Redis
	redisClient := redis.NewRedisClient(&conf.Connection.Redis)
	defer redisClient.Close()
	appLogger.Info("Redis connected")

	//* ====================== Running Server ======================

	server.Run(conf, &dbList, appLogger)
}
