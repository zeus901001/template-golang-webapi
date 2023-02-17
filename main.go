package main

import (
	"golang-webapi/controllers/account"
	"golang-webapi/controllers/auth"
	"golang-webapi/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	mysql.Connect()
	// mysql.Migrate()

	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.107.209"})

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Web API Http Server is running...")
	})

	api := router.Group("/api")
	{
		api.POST("/auth/login", auth.Login)
		api.POST("/auth/register", auth.Register)
		api.POST("/auth/updateAccessToken", auth.UpdateAccessToken)
		api.POST("/auth/removeRefreshToken", auth.RemoveRefreshToken)

		api.GET("/account/users", account.Users)
		api.GET("/account/users/:id", account.UserEdit)
	}

	router.Run("localhost:8080")
}
