package models

import "time"

// Comment is struct hold unit of request and response
type Comment struct {
	ID      int64
	Name    string
	Text    string
	Created time.Time
	Updated time.Time
}

// PreInsert create and update field.
func (c *Comment) PreInsert() error { return nil }

// PreUpdate update field.
func (c *Comment) PreUpdate() error { return nil }

type Cat struct {
	ID   int64
	Name string
	Type string
}

func (c *Cat) PreInsert() error { return nil }

func (c *Cat) ReadOne(fields string) {
	_ = DB().Select(fields).First(c, c.ID)
}

func (c *Cat) ReadMore(sort, fields string, page, perPage int64) (cats []Cat) {
	DB().Select(fields).Order(sort).Offset(page * perPage).Limit(perPage).Find(&cats)
	return nil
}

type HealthCheck struct {
	Message string `json:"message"`
}
