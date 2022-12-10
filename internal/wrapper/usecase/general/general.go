package general

import (
	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/repository"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type GeneralUsecase struct {
}

func NewGeneralUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) GeneralUsecase {
	return GeneralUsecase{}
}
