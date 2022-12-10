package cms

import (
	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type CMSRepository struct {
}

func NewCMSRepository(conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CMSRepository {
	return CMSRepository{}
}
