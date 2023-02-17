package auth

import (
	"golang-webapi/models"

	"github.com/gin-gonic/gin"
)

/* Update access token using refresh token. */
func UpdateAccessToken(context *gin.Context) {

}

/* Remove expired refresh token. */
func RemoveRefreshToken(context *gin.Context) {

}

/* Generate access token for user. */
func GenerateAccessTokenForUser(user *models.User) string {
	return ""
}

/* Generate refresh token for user. */
func GenerateRefreshTokenForUser() string {
	return ""
}
