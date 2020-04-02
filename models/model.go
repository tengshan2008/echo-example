package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Comment is struct hold unit of request and response
type Comment struct {
	ID        uint       `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"update_time,omitempty"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Content   string     `json:"content,omitempty"`
	Author    string     `json:"author,omitempty"`
	CatID     uint       `json:"cat_id,omitempty"`
}

func contain(v string) string {
	return fmt.Sprintf("%%%s%%", v)
}

// Insert one comment
func (c *Comment) Insert() bool {
	db := DB().Create(c)
	return db.NewRecord(c)
}

// ReadOne get one comment record
func (c *Comment) ReadOne(fields string) {
	var db *gorm.DB
	if fields != "" {
		db = DB().Select(fields)
	} else {
		db = DB()
	}
	_ = db.First(c, c.ID)
}

// ReadMore get multiple comment record
func (c *Comment) ReadMore(search, sort, fields string, page, perPage int64) (comments []Comment) {
	var db *gorm.DB
	if fields != "" {
		db = DB().Select(fields)
	} else {
		db = DB()
	}
	if perPage == 0 {
		perPage = 20
	}
	if c.Author != "" {
		db = db.Where("author LIKE ?", contain(c.Author))
	}
	if c.Content != "" {
		db = db.Where("content LIKE ?", contain(c.Content))
	}
	if search != "" {
		db = db.
			Where("author LIKE ?", contain(search)).
			Or("content LIKE ?", contain(search))
	}
	_ = db.Where("cat_id", c.CatID).Order(sort).Offset(page * perPage).Limit(perPage).Find(&comments)
	return
}

// Delete remove one comment
func (c *Comment) Delete() bool {
	db := DB().Delete(c)
	return db.NewRecord(c)
}

// Cat table
type Cat struct {
	ID        uint       `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Name      string     `json:"name,omitempty"`
	Type      string     `json:"type,omitempty"`
}

// Insert one cat record
func (c *Cat) Insert() bool {
	db := DB().Create(c)
	return db.NewRecord(c)
}

// ReadOne get one cat record
func (c *Cat) ReadOne(fields string) {
	var db *gorm.DB
	if fields != "" {
		db = DB().Select(fields)
	} else {
		db = DB()
	}
	_ = db.First(c, c.ID)
}

// ReadMore get mutiple cat record
func (c *Cat) ReadMore(search, sort, fields string, page, perPage int64) (cats []Cat) {
	var db *gorm.DB
	if fields != "" {
		db = DB().Select(fields)
	} else {
		db = DB()
	}
	if perPage == 0 {
		perPage = 20
	}
	if c.Name != "" {
		db = db.Where("name LIKE ?", contain(c.Name))
	}
	if c.Type != "" {
		db = db.Where("type LIKE ?", contain(c.Type))
	}
	if search != "" {
		db = db.
			Where("name LIKE ?", contain(search)).
			Or("type LIKE ?", contain(search))
	}
	_ = db.Order(sort).Offset(page * perPage).Limit(perPage).Find(&cats)
	return
}

// Delete one cat record
func (c *Cat) Delete() bool {
	db := DB().Delete(c)
	return db.NewRecord(c)
}

// HealthCheck message
type HealthCheck struct {
	Message string `json:"message"`
}
