package controllers

import (
	"github.com/gin-gonic/gin"
	"library/models"
	"library/services"
	"net/http"
	"strconv"
)

type BookController struct {
	bookService services.BookService
}

func NewBookController(bookService services.BookService) *BookController {
	return &BookController{bookService: bookService}
}

// CreateBook 创建图书
func (b *BookController) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := b.bookService.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": book})
}

// GetBookByID 根据ID获取图书
func (b *BookController) GetBookByID(c *gin.Context) {
	id_t, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(id_t)
	book, err := b.bookService.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook 更新图书
func (b *BookController) UpdateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := b.bookService.UpdateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook 删除图书
func (b *BookController) DeleteBook(c *gin.Context) {
	id_t, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(id_t)
	if err := b.bookService.DeleteBook(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
