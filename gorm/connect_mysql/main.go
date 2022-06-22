package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	// Code  sql.NullString
	Code string
	Price uint
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

	// 定义表结构 直接生成对应的表 migration
	// 初始化表的时候使用 也就是建表
	// _ = db.AutoMigrate(&Product{})

	// 进行增删查改等操作
	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)

	// Update - 更新多个字段
	// 注意有 Update 和 Updates 两个方法
	// 仅更新非零值字段 例如下面的Code为空 也就是字符串类型的零值 所以不更新 Code 字段 只更新 Price 字段
	// db.Model(&product).Updates(Product{Price: 200, Code: ""})
	// 如果想设置零值 例如字符串零值 可以在定义 struct 时使用 NullString 类型
	// db.Model(&product).Updates(Product{Price: 200, Code: sql.NullString{String: "F42", Valid: true}})
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product 并没有执行 delete 语句 是软删除
	db.Delete(&product, 1)
}
