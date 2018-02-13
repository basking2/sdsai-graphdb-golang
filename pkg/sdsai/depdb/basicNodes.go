package depdb

// This file is concerned with reading and writing basic type nodes.

import (
	"github.com/basking2/sdsai-graphdb-golang/pkg/sdsai/graphdb"
	"time"
)

var basic_key_prefix = "basic_"

func basic2key(name string) string {
	return basic_key_prefix + graphdb.ScrambleKey(name)
}

// Root nodes are loaded and generally retained in memory.
func (db *DepDb) NewBasicNode(name string, timeout int64) (graphdb.NodeId, error) {
	data := DepData{}
	data.Name = name
	depId, err := db.graphdb.Put(basic2key(name), &data, nil, nil)
	if err != nil {
		return "", err
	}

	if timeout > 0 {
		db.timeouts[string(depId)] = time.Now().Unix() + timeout
	}

	return depId, err
}

func (db *DepDb) FindAllProducts() ([]graphdb.NodeId, error) {
	ids, err := db.graphdb.FindIds(basic_key_prefix)
	if err != nil {
		return make([]graphdb.NodeId, 0), err
	}

	deps := make([]graphdb.NodeId, len(ids))

	for i, id := range ids {
		deps[i] = id
	}

	return deps, nil
}
