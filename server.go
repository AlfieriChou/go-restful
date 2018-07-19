package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/fvbock/endless"
	controller "gin-demo/controller"
	router "gin-demo/router"
)

func main() {
	handler := router.Init()
	endless.ListenAndServe(":8000", handler)
}