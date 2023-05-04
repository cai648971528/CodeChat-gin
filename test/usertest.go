package main

import (
	"chat-gin/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:root@tcp(10.64.5.206:31234)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.UserBasic{})

	//data := model.UserBasic{Name: "test1", PassWord: "123456", Phone: "17888889999"}
	//db.Create(&data)

	//var user model.UserBasic
	//fmt.Println(db.First(&user))
}
