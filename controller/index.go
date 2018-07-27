package controller

import (
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
)

func IndexGET (c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}