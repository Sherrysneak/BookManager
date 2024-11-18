package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
	ctx = context.Background()
)

func InitDB() {
	dsn := "sherry_user:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败!")
	}
	fmt.Println("数据库连接成功!")
}

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func GetDB() *gorm.DB {
	return DB
}

// 关闭数据库连接
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("获取 sqlDB 实例失败!")
		return
	}

	err = sqlDB.Close() // 关闭 sql.DB 连接
	if err != nil {
		fmt.Println("关闭数据库连接失败:", err)
	} else {
		fmt.Println("数据库连接已关闭.")
	}
}
