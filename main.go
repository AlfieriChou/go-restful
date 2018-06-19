package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {

	db, err := gorm.Open("mysql", "root:@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")

    if err != nil {
        fmt.Println(err)
        panic("--- 数据库连接失败")
    }
    defer db.Close()

	db.AutoMigrate(&Person{})

	r := gin.Default()
	r.GET("/", GetProjects)

	r.Run(":8080")
}

func GetProjects(c *gin.Context) {
	var person []Person
	if err := db.Find(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}