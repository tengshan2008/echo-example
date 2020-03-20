package models

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

const DbPath = "/tmp/gorm.db"

var db *gorm.DB

func New() *gorm.DB {
	log.Debug("Model NewDB")

	newDb, err := newDB(DbPath)
	if err != nil {
		panic(err)
	}
	db = migrate(newDb)
	return db
}

func newDB(path string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", path)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db, errors.Wrap(err, "open database failed")
}

func DB() *gorm.DB {
	return db
}

func migrate(db *gorm.DB) *gorm.DB {
	return db.AutoMigrate(new(Cat), new(Comment))
}
