package models

import (
	"github.com/labstack/gommon/log"

	"github.com/jinzhu/gorm"
	// sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
)

// DbPath config
const DbPath = "/tmp/gorm.db"

var db *gorm.DB

// DB get database object
func DB() *gorm.DB {
	return db
}

// ConnectDB connect db
func ConnectDB() *gorm.DB {
	log.Info("model connect database")

	if err := connectDB(DbPath); err != nil {
		log.Errorf("connect database failed: %v", err)
		panic(err)
	}

	db = migrate()
	return db
}

func connectDB(path string) (err error) {
	db, err = gorm.Open("sqlite3", path)

	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return errors.Wrap(err, "open database failed")
}

func migrate() *gorm.DB {
	return db.AutoMigrate(new(Cat), new(Comment))
}
