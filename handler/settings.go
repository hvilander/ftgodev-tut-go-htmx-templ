package handler 

import (
  "fmt"
  "net/http"
  "ftgodev-tut/db"
  "ftgodev-tut/view/settings"
  "ftgodev-tut/pkg/kit/validate"
)


func HandleSettingsIndex(w http.ResponseWriter, r *http.Request) error {
  // make sure user is authed
  user := getAuthenticatedUser(r)

  return render(r, w, settings.Index(user))

}

func HandleProfile(w http.ResponseWriter, r *http.Request) error {
  params := settings.ProfileParams{
    Username: r.FormValue("username"),
  }
  fmt.Println("HELO from Profile update")
  errors := settings.ProfileFormErrors{}

  ok := validate.New(&params, validate.Fields{
    "Username": validate.Rules(validate.Min(2), validate.Max(50)),
  }).Validate(&errors)

  if !ok {
    return render(r, w, settings.ProfileSettings(params, errors))
  }
  user := getAuthenticatedUser(r)
  user.Account.Username = params.Username

  if err := db.UpdateAccount(&user.Account); err != nil {
    return err
  }
  params.Success = true;

  return render(r, w, settings.ProfileSettings(params, settings.ProfileFormErrors{}))
}
