package main

import (
  "github.com/BurntSushi/toml"
)


func main() {
  c := config{}

  toml.Decode(`
    [Storage]
    Type = "local"
    Path = "/tmp/graphdb"
    `, &c)

  println("Got Type: ", c.Storage.Type)
  println("Got Path: ", c.Storage.Path)
  serve(&c)
}
