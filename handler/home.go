package handler

import (
  "net/http"

  "ftgodev-tut/view/home"
)


func HandleHomeIndex(w http.ResponseWriter, r *http.Request) {

  home.Index().Render(r.Context(), w)

}
