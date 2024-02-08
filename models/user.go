package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	id          uint `gorm:primaryKey; not null"`
	login       string
	code        string
	is_writer   bool
	is_executor bool
	is_admin    bool
}
