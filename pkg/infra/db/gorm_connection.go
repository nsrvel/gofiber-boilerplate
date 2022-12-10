package db

import (
	"time"

	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DatabaseList struct {
	SimpleTransaction *gorm.DB
}

func NewGORMConnection(conf *config.DatabaseAccount, log *logrus.Logger) *gorm.DB {

	var db *gorm.DB
	var err error

	//* Get DBName from DB source
	dbName := utils.GetDBNameFromDriverSource(conf.DriverSource)

	//* GORM Configuration
	gormConf := &gorm.Config{
		//* Disable gorm log
		Logger: gormlog.Default.LogMode(gormlog.LogLevel(gormlog.Error)),
		// Logger: gormlog.Default.LogMode(gormlog.LogLevel(gormlog.Info)),
		// Logger: gormlog.Default.LogMode(gormlog.Silent),
		//* Table name is singular
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		//* Skip default gorm tx
		// SkipDefaultTransaction: true,
	}

	//* Open Connection depend on driver
	if conf.DriverName == "postgres" || conf.DriverName == "pgx" {
		db, err = gorm.Open(postgres.Open(conf.DriverSource), gormConf)
	}
	if conf.DriverName == "sqlserver" {
		db, err = gorm.Open(sqlserver.Open(conf.DriverSource), gormConf)
	}

	if err != nil {
		log.Fatal("Failed to connect database " + dbName + ", err: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to connect database " + dbName + ", err: " + err.Error())
	}
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDB.SetConnMaxIdleTime(conf.ConnMaxIdleTime * time.Minute)
	sqlDB.SetConnMaxLifetime(conf.ConnMaxLifetime * time.Minute)

	log.Info("Connection Opened to Database " + dbName)
	return db
}
