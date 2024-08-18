package handler

import (
  "fmt"
  "net/http"

 // "ftgodev-tut/db"
  "ftgodev-tut/view/home"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
  user := getAuthenticatedUser(r)
  fmt.Printf("%+v\n", user.Account) // prints the field name and value of a struct 

  return home.Index().Render(r.Context(), w)
}
