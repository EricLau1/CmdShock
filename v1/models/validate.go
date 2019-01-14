package models

import "fmt"

func Unique(table, attr, value string) (int, error) {

  con := Connect()

  sql := fmt.Sprintf("select count(*) from %s where %s = ?", table, attr)

  var count int

  err := con.QueryRow(sql, value).Scan(&count)

  if err != nil {
  
    return -1, err

  }

  defer con.Close()

  return count, nil
}

func IsEmpty(command Command) (bool) {
  
  if command.Name == "" || command.Description == "" {
    
    return true

  }

  return false

}
