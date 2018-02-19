package main

import (
  "net/http"
)


func serve(c *config) error {
  s := http.Server{}
  s.Addr = "0.0.0.0:8080"
  return s.ListenAndServe()
}
