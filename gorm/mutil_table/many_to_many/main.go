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

// User has and belongs to many languages, `user_languages` is the join table
type UserLan struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
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

	// db.AutoMigrate(&UserLan{})

	// add recode
	//lan := []Language{}
	//lan = append(lan, Language{Name: "go"}, Language{Name: "Python"})
	//user := UserLan{
	//	Languages: lan,
	//}
	//
	//db.Create(&user)

	// 查询记录
	var userLan UserLan
	db.Preload("Languages").First(&userLan)
	for _, lan := range userLan.Languages {
		fmt.Println(lan.Name)
	}


}
