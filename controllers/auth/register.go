package auth

import (
	"golang-webapi/models"
	"golang-webapi/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* Register action. */
func Register(c *gin.Context) {
	var user models.User

	rec := mysql.DB.First(&user, c.Params.ByName("email"))
	if rec.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if rec.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"success": false, "msg": "Your email already exists."})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := mysql.DB.Create(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	accessToken := GenerateAccessTokenForUser(&user)
	refreshToken := GenerateRefreshTokenForUser()

	c.JSON(http.StatusCreated, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}
