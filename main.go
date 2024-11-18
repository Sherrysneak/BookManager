// main.go

package main

import (
	"library/config"
	"library/models"
	"library/routes"
)

func main() {
	// 初始化数据库
	config.InitDB()
	defer config.CloseDB()

	db := config.GetDB()
	err := db.AutoMigrate(&models.Book{})
	err = db.AutoMigrate(&models.User{})
	err = db.AutoMigrate(&models.Borrowing{})
	if err != nil {
		panic("迁移数据库失败: " + err.Error())
	}
	// 设置路由
	r := routes.SetupRouter()
	r.Run()
}
