package routes

import (
	"crud-golang/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/", controllers.GetUser)
	r.GET("/create", controllers.CreatePage)
	r.POST("/", controllers.CreateUser)
	r.GET("/update/:id", controllers.UpdatePage)
	r.POST("/update/:id", controllers.UpdateUser)
	r.GET("/delete/:id", controllers.DeleteUser)
}
