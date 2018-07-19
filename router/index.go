package router

import (
	"gopkg.in/gin-gonic/gin.v1"
	controller "gin-demo/controller"
)

func Init () *gin.Engine {
	r := gin.New()
	api := r.Group("")
	api.GET("/", controller.IndexGET)
	return r
}
