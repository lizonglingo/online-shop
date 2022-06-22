package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string	// 使用指针和null string都可以解决空字符串的问题
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
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

	// 通过First查询单个数据
	// 通过主键升序
	//var user User
	//result := db.First(&user)
	//is := errors.Is(result.Error, gorm.ErrRecordNotFound)
	//if is {
	//	fmt.Println("data not found")
	//}
	//fmt.Println(user.ID)

	// 通过主键查询 看文档

	// 检索全部对象
	//var users []User
	//result := db.Find(&users)
	//fmt.Printf("%d lines\n", result.RowsAffected)
	//for _, user := range users {
	//	fmt.Println(user)
	//}

	// 进行条件查询 看文档
	var user User
	db.Where("name = ?", "jinzhu1").First(&user)
	// 下面这种方式屏蔽了底层细节 不用再记住具体的列明叫什么
	// db.Where(&User{Name: "jinzhu1"}).First(&user)
	fmt.Println(user)
}
