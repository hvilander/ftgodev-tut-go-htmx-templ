package handler

import (
  "fmt"
  "net/http"
  //"github.com/replicate/replicate-go"
  "encoding/json"
)

type ReplicateResp struct {
  Status string `json:"status"`
  Output []string `json:"output"`
}

func HandleReplicateCallback(w http.ResponseWriter, r *http.Request) error {
  var resp ReplicateResp 
  if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
    return err
  }


  fmt.Println(resp)

  return nil
}
