package controllers

import (
	"crud-go/config"
	"crud-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := config.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var existingUser models.User
	if err := config.DB.Where("name = ? OR email = ?", user.Name, user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "Email already exists"})
		return
	}
	config.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	var input models.User
	c.ShouldBindJSON(&input)
	user.Name = input.Name
	user.Email = input.Email
	config.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.User{}, "id = ?", id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
