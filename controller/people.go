package controller

import (
  "net/http"
  "strconv"
  . "go-restful/models"
  "gopkg.in/gin-gonic/gin.v1"
)

func GetPeople(c *gin.Context) {

  tmpid := c.Param("id")
  id, err := strconv.Atoi(tmpid)
  if err != nil {
    c.JSON(http.StatusOK, gin.H{
      "success": false, "msg": "格式错误",
    })
  } else {
    people := GetPeopleById(id)
    if people == nil{
      c.JSON(http.StatusOK, gin.H{
        "success": true, "data": "",
      })
    }else{
      c.JSON(http.StatusOK, gin.H{
        "success": true, "data": people,
      })
    }
  }
}