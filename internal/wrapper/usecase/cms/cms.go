package cms

import (
	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/repository"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type CMSUsecase struct {
}

func NewCMSUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CMSUsecase {
	return CMSUsecase{}
}
