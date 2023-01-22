package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Conect() {
	d, err := gorm.Open("mysql", "postgres:a1980/test?charset=utf8&pqrceTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
