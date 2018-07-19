package database

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db gorm.DB

func Database() {
    db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
    if err != nil {
      panic("failed to connect database")
    }
    defer db.Close()
}