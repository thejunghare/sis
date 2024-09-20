package controllers

import (
	"erp/api/api/models"
	"erp/api/config"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var db *gorm.DB = config.Connectdb()

var response struct {
	Password string
	Id       int
}

func Login(c *gin.Context) {
	var user models.User
	status := 0

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	if err := db.Raw("SELECT password, id FROM users WHERE email = ? ", user.Username).Scan(&response).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})
		return
	}

	if user.Password == response.Password {
		// for debugging
		status = 1
		fmt.Println(status)
		fmt.Println(response.Id)

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": user.ID,
		})

		tokenString, err := token.SignedString()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "falied to create token",
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "wrong password or email!",
		})
	}

}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
}
