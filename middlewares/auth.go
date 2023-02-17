package middlewares

import "github.com/gin-gonic/gin"

/* Verify your access token. */
func VerifyAccessToken(accessToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

/* Verify your refresh token. */
func VerifyRefreshToken(refreshToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
