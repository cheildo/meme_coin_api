package controllers

import (
	"net/http"

	"github.com/cheildo/meme_coin_api/api/models"
	"github.com/cheildo/meme_coin_api/api/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateMemeCoin(c *gin.Context) {
	var coin models.MemeCoin
	if err := c.ShouldBindJSON(&coin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	err := validate.Struct(coin)
	if err != nil {
		// Extract validation errors
		validationErrors := err.(validator.ValidationErrors)
		errorsMap := make(map[string]string)
		for _, fieldErr := range validationErrors {
			errorsMap[fieldErr.Field()] = fieldErr.Tag()
		}
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": errorsMap})
		return
	}

	newCoin, err := services.CreateMemeCoin(coin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newCoin)
}

func GetMemeCoin(c *gin.Context) {
	id := c.Param("id")
	coin, err := services.GetMemeCoin(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meme coin not found"})
		return
	}
	c.JSON(http.StatusOK, coin)
}

func UpdateDescription(c *gin.Context) {
	id := c.Param("id")
	var payload struct {
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.UpdateDescription(id, payload.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Description updated"})
}

func DeleteMemeCoin(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteMemeCoin(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Meme coin deleted"})
}

func PokeMemeCoin(c *gin.Context) {
	id := c.Param("id")
	err := services.PokeMemeCoin(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Meme coin poked"})
}
