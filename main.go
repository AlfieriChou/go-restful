package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
  	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, _ := gorm.Open("mysql", "root:@test?charset=utf-8&parseTime=True&loc=Local")
	defer db.Close()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.Run()
}