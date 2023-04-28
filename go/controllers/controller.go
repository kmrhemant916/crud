package controllers

import (
	"crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserInput struct {
	ID uuid.UUID `gorm:"type:char(36);primary_key;"`
	Email  string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{ID: uuid.New(), Email: input.Email, Password: input.Password}
	db := c.MustGet("db").(*gorm.DB)
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exist"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": input})
}