package controllers

import (
	"erp/api/api/models"
	"erp/api/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = config.Connectdb()

func Login(c *gin.Context) {
	var user models.User
	var password string
	status := 0

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	if err := db.Raw("SELECT password FROM users WHERE email = ? ", user.Username).Scan(&password).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})
		return
	}

	if user.Password == password {
		status = 1
		fmt.Println(status)
	} else {
		fmt.Println(status)
	}

}
