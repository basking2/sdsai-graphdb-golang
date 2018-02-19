package main

import (
  "github.com/BurntSushi/toml"
  "github.com/graymeta/stow"
  _ "github.com/basking2/sdsai-graphdb-golang/pkg/sdsai/graphdb"
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

  st, err := stow.Dial(c.Storage["type"], c.Storage)
  if err != nil {
    panic(err.Error())
  }

  st.Close()

  serve(&c)
}
