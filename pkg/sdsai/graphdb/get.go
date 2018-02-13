package graphdb

import (
	"encoding/json"
	"github.com/graymeta/stow"
	"io"
)

// Using a node id as returned from the storage system, fetch a node.
//
// The Id field in the node struct is related to, but different this value.
func (db *GraphDb) Get(id NodeId) (*Node, error) {
	data, err := db.container.Item(string(id))
	if err != nil {
		return nil, err
	}

	sz, err := data.Size()
	if err != nil {
		return nil, err
	}
	bytes := make([]byte, sz)

	in, err := data.Open()
	if err != nil {
		return nil, err
	}

	defer in.Close()
	io.ReadFull(in, bytes)

	node := Node{}

	if err := json.Unmarshal(bytes, &node); err != nil {
		return nil, err
	}

	if node.InEdges == nil {
		node.InEdges = make([]NodeId, 0)
	}

	if node.OutEdges == nil {
		node.InEdges = make([]NodeId, 0)
	}

	return &node, nil
}

func (db *GraphDb) GetItemData(data stow.Item, userdata interface{}) (*Node, error) {
	sz, err := data.Size()
	if err != nil {
		return nil, err
	}
	bytes := make([]byte, sz)

	in, err := data.Open()
	if err != nil {
		return nil, err
	}

	defer in.Close()
	io.ReadFull(in, bytes)

	node := Node{}
	node.Data = userdata

	if err := json.Unmarshal(bytes, &node); err != nil {
		return nil, err
	}

	if node.InEdges == nil {
		node.InEdges = make([]NodeId, 0)
	}

	if node.OutEdges == nil {
		node.InEdges = make([]NodeId, 0)
	}

	return &node, nil
}

func (db *GraphDb) GetData(id NodeId, userdata interface{}) (*Node, error) {
	data, err := db.container.Item(string(id))
	if err != nil {
		return nil, err
	}

	return db.GetItemData(data, userdata)
}
