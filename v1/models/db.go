package models

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

const DRIVER = "mysql"
const DBNAME = "cmdshock"
const USER   = "root"
const PASS   = "@root"

func Connect() *sql.DB {

  URL := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", USER, PASS, DBNAME)

  con, err := sql.Open(DRIVER, URL)

  if err != nil {
  
    panic(err.Error())

  }

  return con
}
