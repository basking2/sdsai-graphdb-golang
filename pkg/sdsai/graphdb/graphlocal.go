package graphdb

import (
  "github.com/graymeta/stow"
  "github.com/graymeta/stow/local"
)

func NewLocalGraphDb(name, path string) (GraphDb, error) {
  config := make(stow.ConfigMap)
	config[local.ConfigKeyPath] = path

	location, err := stow.Dial(local.Kind, config)
	if err != nil {
		return GraphDb{}, err
	}

  return NewGraphDb(name, location)

}
