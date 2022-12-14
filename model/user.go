package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
	Password   string `json:"-"`
	Token      string `json:"-"`
}
