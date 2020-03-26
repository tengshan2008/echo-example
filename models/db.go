package models

import (
	"github.com/labstack/gommon/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
)

const DbPath = "/tmp/gorm.db"

var db *gorm.DB

func New() *gorm.DB {
	log.Debug("Model NewDB")

	err := newDB(DbPath)
	if err != nil {
		panic(err)
	}
	db = migrate()
	return db
}

func newDB(path string) (err error) {
	db, err = gorm.Open("sqlite3", path)

	// db.DB().SetMaxIdleConns(10)
	// db.DB().SetMaxOpenConns(100)

	return errors.Wrap(err, "open database failed")
}

func DB() *gorm.DB {
	return db
}

func migrate() *gorm.DB {
	return db.AutoMigrate(new(Cat), new(Comment))
}
