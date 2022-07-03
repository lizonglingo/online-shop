package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type User struct {
	gorm.Model
	// 指定外键
	CreditCards []CreditCard `gorm:"foreignKey:UserID"`
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	// 设置全局logger 打印每一行sql
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.Lshortfile),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// 进行数据库的 配置和连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	//db.AutoMigrate(&User{})
	//db.AutoMigrate(&CreditCard{})

	//user := User{}
	//db.Create(&user)
	//db.Create(&CreditCard{
	//	Number: "12",
	//	UserID: user.ID,
	//})
	//db.Create(&CreditCard{
	//	Number: "34",
	//	UserID: user.ID,
	//})

	var user User
	db.Preload("CreditCards").First(&user)
	for _, card := range user.CreditCards {
		fmt.Println(card.Number)
	}

}
