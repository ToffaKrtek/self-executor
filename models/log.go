package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	id         uint `gorm:primaryKey; not null"`
	text       string
	user_id    uint
	script_id  uint
	created_at string
	ip_address string
}
