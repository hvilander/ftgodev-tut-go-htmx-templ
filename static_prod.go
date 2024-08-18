//go:build !dev
// +build ~!dev
package main

import (
  "fmt"
  "embed"
  "net/http"
)

//go:embed public
var publicFS embed.FS

func public() http.Handler {
  fmt.Println("building static files for PROD")
  return http.FileServerFS(publicFS)
}
