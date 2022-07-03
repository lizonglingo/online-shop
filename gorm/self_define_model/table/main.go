package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type Language struct {
	gorm.Model
	Name string
}

// 实现接口自定义表名
//func (Language) TableName() string {
//	return "my_language"
//}

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
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "duo_",		// 自定义表前缀 注意这个配置不能和自定义表名一起使用
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Language{})



}
