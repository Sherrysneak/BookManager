// routes/routes.go

package routes

import (
	"github.com/gin-gonic/gin"
	"library/config"
	"library/controllers"
	"library/repositories"
	"library/services"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 设置数据库连接
	db := config.GetDB()

	// 设置 BookController
	bookRepo := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	bookController := controllers.NewBookController(*bookService)

	// 设置 UserController
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(*userService)

	// 设置 BorrowingController
	borrowRepo := repositories.NewBorrowingRepository(db)
	borrowService := services.NewBorrowingService(borrowRepo, bookRepo)
	borrowController := controllers.NewBorrowingController(*borrowService)

	// Book 路由
	r.POST("/books", bookController.CreateBook)
	r.GET("/books/:id", bookController.GetBookByID)
	r.PUT("/books/:id", bookController.UpdateBook)
	r.DELETE("/books/:id", bookController.DeleteBook)

	// User 路由
	r.POST("/users", userController.CreateUser)
	r.GET("/users/:id", userController.GetUserByID)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	// Borrowing 路由
	r.POST("/borrow", borrowController.BorrowBook)
	r.POST("/return", borrowController.ReturnBook)
	r.GET("/borrowings/:user_id", borrowController.GetUserBorrowings)

	return r
}
