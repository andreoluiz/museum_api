package controllers

import (
	"io/ioutil"
	"mime/multipart"
	"museum-api/database"
	"museum-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateMuseum(c *gin.Context) {
	var museum models.Museum

	// Bind form data
	museum.Title = c.PostForm("title")
	museum.Description = c.PostForm("description")
	museum.Category1 = c.PostForm("category1")
	museum.Category2 = c.PostForm("category2")
	museum.Link = c.PostForm("link")
	museum.Address = c.PostForm("address")
	museum.CEP = c.PostForm("cep")
	museum.City = c.PostForm("city")
	museum.State = c.PostForm("state")
	museum.Information = c.PostForm("information")

	// Decode image
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get image"})
		return
	}

	imageFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open image"})
		return
	}
	defer func(imageFile multipart.File) {
		err := imageFile.Close()
		if err != nil {

		}
	}(imageFile)

	imageBytes, err := ioutil.ReadAll(imageFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image"})
		return
	}
	museum.Image = imageBytes

	// Pega o ID do gerente pelo token
	managerID, exists := c.Get("manager_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Manager ID not found in token"})
		return
	}
	museum.ManagerID = managerID.(uint)

	// Salva museu no banco de dados
	if err := database.DB.Create(&museum).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Museu cadastrado com sucesso", "museum": museum})
}

// UpdateMuseum atualiza um museu no banco de dados, alterando apenas
// os campos que foram passados no corpo da requisição
func UpdateMuseum(c *gin.Context) {
	var updateData models.Museum
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	museumID := c.Param("id")
	var museum models.Museum
	if err := database.DB.Where("id = ?", museumID).First(&museum).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Museum not found"})
		return
	}

	if updateData.Title != "" {
		museum.Title = updateData.Title
	}
	if updateData.Description != "" {
		museum.Description = updateData.Description
	}
	if len(updateData.Image) > 0 {
		museum.Image = updateData.Image
	}
	if updateData.Category1 != "" {
		museum.Category1 = updateData.Category1
	}
	if updateData.Category2 != "" {
		museum.Category2 = updateData.Category2
	}
	if updateData.Link != "" {
		museum.Link = updateData.Link
	}
	if updateData.Address != "" {
		museum.Address = updateData.Address
	}
	if updateData.CEP != "" {
		museum.CEP = updateData.CEP
	}
	if updateData.City != "" {
		museum.City = updateData.City
	}
	if updateData.State != "" {
		museum.State = updateData.State
	}
	if updateData.Information != "" {
		museum.Information = updateData.Information
	}

	if err := database.DB.Save(&museum).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Museum updated successfully"})
}

func GetMuseumsByState(c *gin.Context) {
	state := c.Query("state")

	var museums []models.Museum
	if err := database.DB.Where("state = ?", state).Find(&museums).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"museums": museums})
}

func GetMuseumsByCity(c *gin.Context) {
	city := c.Query("city")

	var museums []models.Museum
	if err := database.DB.Where("city = ?", city).Find(&museums).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"museums": museums})
}

func GetMuseumsByName(c *gin.Context) {
	name := c.Query("name")

	var museums []models.Museum
	if err := database.DB.Where("title LIKE ?", "%"+name+"%").Find(&museums).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"museums": museums})
}

func DisableMuseum(c *gin.Context) {
	museumID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid museum ID"})
		return
	}

	var museum models.Museum
	if err := database.DB.First(&museum, museumID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Museum not found"})
		return
	}

	museum.Active = false
	if err := database.DB.Save(&museum).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Museum disabled successfully"})
}
