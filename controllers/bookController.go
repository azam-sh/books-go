package controllers

import (
	"net/http"
	"strconv"

	"github.com/azam-sh/books/initializers"
	"github.com/azam-sh/books/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddBook(c *gin.Context) {

	var body models.RequestBook

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	book := models.Book{Name: body.Name, AccessID: body.AccessID, CategoryID: body.CategoryID}
	result := initializers.DB.Create(&book)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "book created!",
	})
}

func GetAllBooks(c *gin.Context) {

	var books []models.Book

	if s := c.Query("search"); s != "" {
		if err := initializers.DB.Joins("Access").Joins("Category").Where("books.name LIKE ?", "%"+s+"%").Scopes(Paginate(c.Request)).Find(&books).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "can't get books",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": books,
		})
		return
	} else {
		if err := initializers.DB.Joins("Access").Joins("Category").Scopes(Paginate(c.Request)).Find(&books).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "can't get books",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": books,
		})
		return
	}
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	if result := initializers.DB.Joins("Access").Joins("Category").First(&book, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "book not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &book,
	})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var body models.RequestBook

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var book models.Book

	if result := initializers.DB.First(&book, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "could not update book",
		})
		return
	}

	book.Name = body.Name
	book.AccessID = body.AccessID

	initializers.DB.Save(&book)

	c.JSON(http.StatusOK, gin.H{
		"message": "book updated successfully",
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	if result := initializers.DB.Delete(&models.Book{}, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "book not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "book deleted successfully",
	})
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
