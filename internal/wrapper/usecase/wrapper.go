package usecase

import (
	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/repository"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/usecase/cms"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/usecase/core"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/usecase/general"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type Usecase struct {
	General general.GeneralUsecase
	Core    core.CoreUsecase
	CMS     cms.CMSUsecase
}

func NewUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) Usecase {
	return Usecase{
		General: general.NewGeneralUsecase(repo, conf, dbList, log),
		Core:    core.NewCoreUsecase(repo, conf, dbList, log),
		CMS:     cms.NewCMSUsecase(repo, conf, dbList, log),
	}
}
