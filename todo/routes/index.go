package routes

import (
	"signup/controller"

	"github.com/gin-gonic/gin"
)

func AppRoutes(r *gin.Engine) {
	r.Static("/register", "./signup")
	r.Static("/login", "./login")
	r.Static("/todo", "./todo")
	r.POST("/signup", controller.Signup)
	r.POST("/loginn", controller.Loginn)
	r.POST("/add",controller.Add)
	r.POST("/getalltodo",controller.GetAllTODO)
	r.DELETE("/delete",controller.Deletetodo)
}
