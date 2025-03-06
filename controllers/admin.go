package controllers

import (
	"log"
	"net/http"
	"strconv"
	"user-management-app/config"
	"user-management-app/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	var response []gin.H
	for _, user := range users {
		response = append(response, gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		})
	}

	c.JSON(http.StatusOK, gin.H{"users": response})
}

func UpdateUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	adminRole, _ := c.Get("role")
	if adminRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can update roles"})
		return
	}

	var input struct {
		Role string `json:"role" binding:"required,oneof=user admin"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.Role == "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot change the role of another admin"})
		return
	}

	user.Role = input.Role
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User role updated successfully"})
}

func DeleteUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Println("Invalid user ID:", userIDStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		log.Println("User not found:", userID)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.Role == "admin" {
		log.Println("Attempt to delete admin prevented")
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot delete an admin"})
		return
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		log.Println("Failed to delete user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	log.Println("User deleted successfully:", userID)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
