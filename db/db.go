package db

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func Init() {
	tmp, err := gorm.Open("mysql", "myweb:myweb@tcp(127.0.0.1:3306)/myweb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	DB = tmp
}

func Close() {
	DB.Close()
}
