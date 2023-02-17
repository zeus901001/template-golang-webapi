package account

import (
	"golang-webapi/models"
	"golang-webapi/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* Find all users. */
func Users(c *gin.Context) {
	var users []models.User

	if err := mysql.DB.Find(&users).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, users)
}

/* Edit user by id. */
func UserEdit(c *gin.Context) {
	var user models.User
	if err := mysql.DB.First(&user, c.Params.ByName("id")).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

/* Remove user by id. */
func UserRemove(c *gin.Context) {
	var user models.User
	if err := mysql.DB.Where("id = ?", c.Params.ByName("id")).Delete(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func UserChangeRole(c *gin.Context) {

}

func UserChangePermission(c *gin.Context) {

}
