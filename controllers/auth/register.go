package auth

import (
	"golang-webapi/models"
	"golang-webapi/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* Register action. */
func Register(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		context.Abort()
		return
	}

	record := mysql.DB.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": record.Error.Error()})
		context.Abort()
		return
	}

	accessToken := GenerateAccessTokenForUser(user)
	refreshToken := GenerateRefreshTokenForUser()

	context.JSON(http.StatusCreated, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}
