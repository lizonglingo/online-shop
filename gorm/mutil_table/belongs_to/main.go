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

// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name string

	// 这里就需要注意去指明外键和所需要连接的表
	CompanyID int // 数据库中被存储的外键字段 company_id
	Company   Company
}

type Company struct {
	ID   int
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

	// db.AutoMigrate(&User{})		// 会自动的也将 company 表进行创建

	//db.Create(&User{
	//	Name: "lzl",
	//	Company: Company{
	//		Name: "mooc",
	//	},
	//})

	//db.Create(&User{
	//	Name: "lzl-2",
	//	Company: Company{
	//		ID: 1},
	//})

	// 使用 Preload 或者 join 进行预加载 完成关联查询
	// 这两种方法生成的sql语句是不同的
	var user User
	db.Preload("Company").First(&user)
	fmt.Println(user.Company)

	db.Joins("Company").First(&user)
	fmt.Println(user.CompanyID)

}
