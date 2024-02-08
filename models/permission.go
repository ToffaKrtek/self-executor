package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	id         uint `gorm:primaryKey; not null"`
	script_id  uint
	user_id    uint
	executable bool
	writable   bool
}
