package server

import (
	"github.com/gin-gonic/gin"
	"sql_multiple1/control"
)

func Init() {
	r := router()
	r.Run("localhost:8080")
}

func router() *gin.Engine {
	r := gin.Default()

	ctrl := control.Controller{}

	r.GET("/user", ctrl.GetAllUsers)
	r.GET("/user/:id", ctrl.GetOneUser)
	r.POST("/user", ctrl.AddUser)
	r.DELETE("/user/:id", ctrl.DeleteUser)

	r.GET("/tag", ctrl.GetAllTags)
	r.GET("/tag/:id", ctrl.GetOneTag)
	r.POST("/tag", ctrl.AddTag)
	r.DELETE("/tag/:id", ctrl.DeleteTag)

	r.POST("/combine", ctrl.UserTagCombine)
	return r
}
