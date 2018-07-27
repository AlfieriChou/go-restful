package router

import (
	"gopkg.in/gin-gonic/gin.v1"
	controller "go-restful/controller"
)

func Init () *gin.Engine {
	r := gin.New()
	api := r.Group("")
	api.GET("/", controller.IndexGET)
	api.GET("/people/:id", controller.GetPeople)
	return r
}
