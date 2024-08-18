//+build dev
//go:build dev
// +build dev

package main

import (
  "fmt"
  "net/http"
  "os"
)

// allows tailwind to rebuid the css without rebuilding the whole binary.
func public() http.Handler {
  fmt.Println("building static files for dev")
  return http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public")))
}
