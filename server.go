package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/fvbock/endless"
	controller "gin-demo/controller"
)

func main() {
	router := gin.Default()
	router.GET("/", controller.IndexGET)
	endless.ListenAndServe(":8000", router)
}