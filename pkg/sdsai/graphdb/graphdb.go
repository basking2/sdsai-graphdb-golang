package graphdb

import (
	"strings"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/op/go-logging"
	"github.com/graymeta/stow"

	// support Azure storage
	_ "github.com/graymeta/stow/azure"
	// support Google storage
	_ "github.com/graymeta/stow/google"
	// support local storage
	_ "github.com/graymeta/stow/local"
	// support swift storage
	_ "github.com/graymeta/stow/swift"
	// support s3 storage
	_ "github.com/graymeta/stow/s3"
	// support oracle storage
	_ "github.com/graymeta/stow/oracle")

type GraphDb struct {
	location  stow.Location
	container stow.Container
	log       *logging.Logger
}

type NodeId string
type NodeData interface{}

type Node struct {
	Name     string   `json:"name"`
	InEdges  []NodeId `json:"inedges"`
	OutEdges []NodeId `json:"outedges"`
	Data     NodeData `json:"data"`
}

// Create a new GraphDb.
func NewGraphDb(name string, location stow.Location) (GraphDb, error) {

	logger := logging.MustGetLogger(name)

	containers, cursor, err := location.Containers(name, stow.CursorStart, pagesize)
	if err != nil {
		return GraphDb{}, err
	}

	for {
		for _, container := range containers {
			if strings.HasSuffix(container.ID(), name) {
				return GraphDb{location, container, logger}, nil
			}
		}

		if stow.IsCursorEnd(cursor) {
			break
		}

		containers, cursor, err = location.Containers(name, cursor, pagesize)
		if err != nil {
			return GraphDb{}, err
		}
	}

	container, err := location.CreateContainer(name)
	if err != nil {
		return GraphDb{}, err
	}

	return GraphDb{location, container, logger}, nil
}

func (db *GraphDb) Close() error {
	return db.location.Close()
}

// Put a node and return the storage ID.
//
// The storage ID is related to, but different than the node's name.
func (db *GraphDb) Put(name string, nodeData interface{}, inEdges, outEdges []NodeId) (NodeId, error) {

	if name == "" {
		return "", errors.New("Node ID may not be empty.")
	}

	node := Node{name, inEdges, outEdges, NodeData(nodeData)}

	return db.putNode(&node)
}

func (db *GraphDb) putNode(node *Node) (NodeId, error) {
	data, err := json.Marshal(node)
	if err != nil {
		return "", err
	}

	reader := bytes.NewReader(data)

	item, err := db.container.Put(node.Name, reader, int64(len(data)), nil)
	if err != nil {
		return "", nil
	}

	return NodeId(item.ID()), nil
}

func (db *GraphDb) Remove(id string) error {
	return db.container.RemoveItem(id)
}
