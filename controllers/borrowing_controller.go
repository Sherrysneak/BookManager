package controllers

import (
	"github.com/gin-gonic/gin"
	"library/services"
	"net/http"
	"strconv"
)

type BorrowingController struct {
	borrowingService services.BorrowingService
}

func NewBorrowingController(borrowingService services.BorrowingService) *BorrowingController {
	return &BorrowingController{borrowingService: borrowingService}
}

// BorrowBook 用户借书的处理器
func (b *BorrowingController) BorrowBook(c *gin.Context) {
	userID, err := strconv.ParseUint(c.PostForm("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	bookID, err := strconv.ParseUint(c.PostForm("book_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	if err := b.borrowingService.BorrowBook(uint(userID), uint(bookID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book borrowed successfully"})
}

// ReturnBook 用户还书的处理器
func (b *BorrowingController) ReturnBook(c *gin.Context) {
	userID, err := strconv.ParseUint(c.PostForm("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	bookID, err := strconv.ParseUint(c.PostForm("book_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	if err := b.borrowingService.ReturnBook(uint(userID), uint(bookID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book returned successfully"})
}

// GetUserBorrowings 获取用户借阅记录的处理器
func (b *BorrowingController) GetUserBorrowings(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	borrowings, err := b.borrowingService.GetUserBorrowings(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, borrowings)
}
