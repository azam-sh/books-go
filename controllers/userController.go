package controllers

import (
	"net/http"

	"books/initializers"
	"books/models"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {

	var users []models.ResponseUser

	initializers.DB.Model(&models.User{}).Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var user models.ResponseUser

	if result := initializers.DB.Model(&models.User{}).First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &user,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var body models.RequestUser

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var user models.User

	if result := initializers.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, result.Error)
		return
	}

	user.FullName = body.FullName
	user.RoleID = body.RoleID
	user.Phone = body.Phone
	user.Active = body.Active

	initializers.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
	})
}
