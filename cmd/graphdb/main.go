package main

import (
  "github.com/BurntSushi/toml"
  "github.com/graymeta/stow"

  // Pull in so all types of storage load.
  "github.com/basking2/sdsai-graphdb-golang/pkg/sdsai/graphdb"
)

func main() {
  c := config{}

  toml.Decode(`
    [Storage]
    type = "local"
    path = "/tmp"
    `, &c)

  println("Got Type: ", c.Storage["type"])
  println("Got Path: ", c.Storage["path"])

  location, err := stow.Dial(c.Storage["type"], c.Storage)
  if err != nil {
    panic(err.Error())
  }

  graphdb, err := graphdb.NewGraphDb("graphdb", location)
  if err != nil {
    panic(err.Error())
  }

  serve(&c)

  graphdb.Close()
}
