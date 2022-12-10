package core

import (
	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type CoreRepository struct {
}

func NewCoreRepository(conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CoreRepository {
	return CoreRepository{}
}
