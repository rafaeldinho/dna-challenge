package sql

import (
	"github/meli/src/shared"
	gormLogger "gorm.io/gorm/logger"

	log "github.com/sirupsen/logrus"
	"github/meli/src/domain/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var logger = log.WithFields(log.Fields{
	"layer": shared.SqlLayer,
})

type SqlRepository interface {
	GetConnection() *gorm.DB
}

type sqlConn struct {
	url string
}

func NewSqlInstance(stringConn string) *sqlConn {
	return &sqlConn{
		url: stringConn,
	}
}

func (r *sqlConn) Migration() {
	migrator := r.GetConnection().Migrator()
	err := migrator.AutoMigrate(
		model.Mutant{},
	)
	if err != nil {
		logger.WithError(err).Error("error executing auto migrate")
	}
}

func (r *sqlConn) GetConnection() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: r.url}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		logger.Error(err.Error())
	}
	return db
}
