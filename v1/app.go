package main

import (
  "fmt"
  "net/http"
  "./routes"
  "./utils"
)

func main(){

  fmt.Println("Listening port 8080")

  utils.LoadTemplates("views/*.html")

  r := routes.NewRouter()

  http.Handle("/", r)

  http.ListenAndServe(":8080", nil)

}

