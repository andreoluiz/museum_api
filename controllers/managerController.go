package controllers

import (
	"museum-api/database"
	"museum-api/models"
	"museum-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateManager(c *gin.Context) {
	var manager models.Manager
	if err := c.ShouldBindJSON(&manager); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(manager.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	manager.Password = hashedPassword

	if err := database.DB.Create(&manager).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Manager created successfully"})
}

func Login(c *gin.Context) {
	var credentials models.ManagerLogin
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var manager models.Manager
	if err := database.DB.Where("email = ?", credentials.Email).First(&manager).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !utils.CheckPasswordHash(credentials.Password, manager.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(manager.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func UpdateManager(c *gin.Context) {
	managerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid manager ID"})
		return
	}

	var manager models.Manager
	if err := database.DB.First(&manager, managerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Manager not found"})
		return
	}

	var newManager models.Manager
	if err := c.ShouldBindJSON(&newManager); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	manager.FirstName = newManager.FirstName
	manager.LastName = newManager.LastName
	manager.Email = newManager.Email

	if err := database.DB.Save(&manager).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Manager updated successfully", "manager": manager})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func DisableManager(c *gin.Context) {
	managerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid manager ID"})
		return
	}

	var manager models.Manager
	if err := database.DB.First(&manager, managerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Manager not found"})
		return
	}

	manager.Active = false
	if err := database.DB.Save(&manager).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Manager account disabled successfully"})
}
