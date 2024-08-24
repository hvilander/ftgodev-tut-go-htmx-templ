package handler

import (
  "fmt"
  "net/http"
  "time"

  "ftgodev-tut/view/home"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
  user := getAuthenticatedUser(r)
  fmt.Printf("%+v\n", user.Account) // prints the field name and value of a struct 

  return home.Index().Render(r.Context(), w)
}


func HandleLongProcess(w http.ResponseWriter, r *http.Request) error {
  time.Sleep(time.Second * 5)
  return home.UserLikes(10000).Render(r.Context(), w) 
}
