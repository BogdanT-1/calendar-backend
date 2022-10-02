package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//containers on the same network cand discover eachother based on name given: "calendar-mysql"
	//can connect network to container even after creation with docker netwrok connect
	d, err := gorm.Open("mysql", "root:password@tcp(calendar-mysql)/calendar?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
