package main

import (
  "net/http"
  "github.com/graymeta/stow"
  "github.com/basking2/sdsai-graphdb-golang/pkg/sdsai/graphdb"
)


func serve(c *config) error {
  location, err := stow.Dial(c.Storage["type"], c.Storage)
  if err != nil {
    panic(err.Error())
  }

  graphdb, err := graphdb.NewGraphDb("graphdb", location)
  if err != nil {
    panic(err.Error())
  }

  mux := http.NewServeMux()
  mux.HandleFunc("/db/", func(rw http.ResponseWriter, req *http.Request){
    // FIXME - write in database manipulation here.
    rw.Write([]byte("Got it.\n"))
  })

  s := http.Server{}
  s.Addr = "0.0.0.0:8080"
  s.RegisterOnShutdown(func() {
    graphdb.Close()
  })
  s.Handler = mux

  return s.ListenAndServe()
}
