package database

import (
  _ "github.com/go-sql-driver/mysql"
  "github.com/go-xorm/xorm"
  "log"
)
var Orm *xorm.Engine
func init() {
  var err error
  Orm, err = xorm.NewEngine("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
  Orm.ShowSQL(true)
  if err != nil {
    log.Fatal(err.Error())
  }
}