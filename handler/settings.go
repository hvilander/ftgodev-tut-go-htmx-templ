package handler 

import (
  "net/http"
  "ftgodev-tut/view/settings"
)


func HandleSettingsIndex(w http.ResponseWriter, r *http.Request) error {
  // make sure user is authed
  user := getAuthenticatedUser(r)

  return render(r, w, settings.Index(user))

}
