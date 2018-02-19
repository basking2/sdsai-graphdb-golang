package main

import (
  "github.com/BurntSushi/toml"
)

func main() {
  c := config{}

  // Load defaults.
  toml.Decode(`
    [Storage]
    type = "local"
    path = "/tmp"
    `, &c)

  println("Got Type: ", c.Storage["type"])
  println("Got Path: ", c.Storage["path"])

  serve(&c)
}
