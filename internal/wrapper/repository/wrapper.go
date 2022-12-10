package repository

import (
	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/repository/cms"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/repository/core"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/repository/general"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	General general.GeneralRepository
	Core    core.CoreRepository
	CMS     cms.CMSRepository
}

func NewRepository(conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) Repository {
	return Repository{
		General: general.NewGeneralRepository(conf, dbList, log),
		Core:    core.NewCoreRepository(conf, dbList, log),
		CMS:     cms.NewCMSRepository(conf, dbList, log),
	}
}
