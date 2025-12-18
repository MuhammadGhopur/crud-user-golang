package controllers

import (
	"crud-golang/config"
	"crud-golang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var users []models.User

	config.DB.Find(&users)

	c.HTML(http.StatusOK, "users.html", gin.H{
		"users": users,
	})
}

func CreatePage(c *gin.Context) {
	c.HTML(http.StatusOK, "create_users.html", nil)
}

func CreateUser(c *gin.Context) {
	var user models.User

	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")

	config.DB.Create(&user)

	c.Redirect(http.StatusFound, "/")
}

func UpdatePage(c *gin.Context) {

	var user models.User

	id, _ := strconv.Atoi(c.Param("id"))

	config.DB.First(&user, id)

	c.HTML(http.StatusOK, "edit_user.html", gin.H{
		"user": user,
	})

}

func UpdateUser(c *gin.Context) {
	var user models.User

	id, _ := strconv.Atoi(c.Param("id"))

	config.DB.First(&user, id)

	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")

	config.DB.Save(&user)

	c.Redirect(http.StatusFound, "/")
}

func DeleteUser(c *gin.Context) {
	var user models.User

	id, _ := strconv.Atoi(c.Param("id"))

	config.DB.First(&user, id)

	config.DB.Delete(&user)

	c.Redirect(http.StatusFound, "/")
}
