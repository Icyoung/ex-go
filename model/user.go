package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Name       string
	Password   string `json:"-"`
	Token      string `json:"-"`
}
