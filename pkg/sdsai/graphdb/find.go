package graphdb

import (
	"errors"
	"strings"
	"github.com/graymeta/stow"
)

var NotFoundError = errors.New("Not found.")

var pagesize = 1000

// Using a node name or ID prefix, find a node.
//
// The value of Id may be used in this call to locate a node.
func (db *GraphDb) Find(idprefix string) ([]*Node, error) {
	items, cursor, err := db.container.Items(idprefix, stow.CursorStart, pagesize)
	if err != nil {
		return nil, err
	}

	// If there is more stuff to get, go get it.
	for !stow.IsCursorEnd(cursor) {
		var items2 []stow.Item
		items2, cursor, err = db.container.Items(idprefix, cursor, pagesize)
		if err != nil {
			return nil, err
		}
		items = append(items, items2...)
	}

	nodes := make([]*Node, len(items))
	for i, v := range items {
		n, err := db.Get(NodeId(v.ID()))
		if err != nil {
			return nil, err
		}
		nodes[i] = n
	}

	return nodes, nil
}

func (db *GraphDb) FindIds(idprefix string) ([]NodeId, error) {
	items, cursor, err := db.container.Items(idprefix, stow.CursorStart, pagesize)
	if err != nil {
		return nil, err
	}

	// If there is more stuff to get, go get it.
	for !stow.IsCursorEnd(cursor) {
		var items2 []stow.Item
		items2, cursor, err = db.container.Items(idprefix, cursor, pagesize)
		if err != nil {
			return nil, err
		}
		items = append(items, items2...)
	}

	nodes := make([]NodeId, len(items))
	for i, v := range items {
		nodes[i] = NodeId(v.ID())
	}

	return nodes, nil
}

func (db *GraphDb) FindOneData(idprefix string, data interface{}) (NodeId, *Node, error) {
	items, cursor, err := db.container.Items(idprefix, stow.CursorStart, pagesize)
	if err != nil {
		return NodeId(""), nil, err
	}

	// If there is more stuff to get, go get it.
	for {
		for _, item := range items {
			if strings.HasSuffix(item.ID(), idprefix) {
				node, err := db.GetItemData(item, data)
				if err != nil {
					return NodeId(""), nil, err
				}

				return NodeId(item.ID()), node, nil
			}
		}

		if stow.IsCursorEnd(cursor) {
			break
		}

		items, cursor, err = db.container.Items(idprefix, cursor, pagesize)
		if err != nil {
			return NodeId(""), nil, err
		}
	}

	return NodeId(""), nil, NotFoundError
}
