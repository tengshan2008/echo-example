package models

import (
	"github.com/jinzhu/gorm"
)

// Comment is struct hold unit of request and response
type Comment struct {
	gorm.Model
	Name string
	Text string
}

// PreInsert create and update field.
func (c *Comment) PreInsert() error { return nil }

// PreUpdate update field.
func (c *Comment) PreUpdate() error { return nil }

type Cat struct {
	gorm.Model
	Name string
	Type string
}

func (c *Cat) Insert() bool {
	return DB().NewRecord(c)
}

func (c *Cat) ReadOne(fields string) {
	_ = DB().Select(fields).First(c, c.ID)
}

func (c *Cat) ReadMore(sort, fields string, page, perPage int64) (cats []Cat) {
	_ = DB().Select(fields).Order(sort).Offset(page * perPage).Limit(perPage).Find(&cats)
	return cats
}

func (c *Cat) Delete() {
	_ = DB().Delete(c)
}

type HealthCheck struct {
	Message string `json:"message"`
}
