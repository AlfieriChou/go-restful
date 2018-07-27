package main

import (
	"github.com/fvbock/endless"
	router "go-restful/router"
)

func main() {
	handler := router.Init()
	endless.ListenAndServe(":8000", handler)
}