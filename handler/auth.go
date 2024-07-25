package handler

import (
  "net/http"
  "ftgodev-tut/view/auth"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
  return auth.Login().Render(r.Context(), w)
}
