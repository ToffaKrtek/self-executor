package models

import "gorm.io/gorm"

type Script struct {
	gorm.Model
	id          uint `gorm:primaryKey; not null"`
	title       string
	description string
	filename    string
	variables   []Variable
}

type Variable struct {
	var_type        string
	var_name        string
	var_description string
}
