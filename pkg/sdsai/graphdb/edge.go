package graphdb

func addedge(ids []NodeId, id NodeId) ([]NodeId, bool) {
	// An empty list is easily defined as a single new edge.
	if ids == nil {
		ids = make([]NodeId, 0, 1)
	}

	// Make sure the edge does not already exist in the list.
	for _, edge := range ids {
		if edge == id {
			return ids, false
		}
	}

	return append(ids, id), true
}

func removeedge(ids []NodeId, id NodeId) ([]NodeId, bool) {
	if ids == nil {
		return make([]NodeId, 0), false
	}

	i := 0
	iend := len(ids)
	for i < iend {
		// If ids[i] == id, replace ids[i] with the end and shorten the list.
		if ids[i] == id {
			iend -= 1
			ids[i] = ids[iend]
		} else {
			// Otherwise, advance the list.
			i += 1
		}
	}

	if iend == len(ids) {
		return ids, false
	} else {
		return ids[0:iend], true
	}
}

// Add an edge going from id1 to id2 nodes.
func (db *GraphDb) AddEdge(id1, id2 NodeId) error {
	node1, err := db.Get(id1)
	if err != nil {
		return err
	}

	node2, err := db.Get(id2)
	if err != nil {
		return err
	}

	var changed1, changed2 bool

	node1.OutEdges, changed1 = addedge(node1.OutEdges, id2)
	node2.InEdges, changed2 = addedge(node2.InEdges, id1)

	if changed1 {
		_, err = db.putNode(node1)
		if err != nil {
			return err
		}
	}

	if changed2 {
		_, err = db.putNode(node2)
		if err != nil {
			return err
		}
	}

	return nil
}

// Remove an edge from id1 to id2.
func (db *GraphDb) RemoveEdge(id1, id2 NodeId) error {
	node1, err := db.Get(id1)
	if err != nil {
		return err
	}

	node2, err := db.Get(id2)
	if err != nil {
		return err
	}

	var changed1, changed2 bool

	node1.OutEdges, changed1 = removeedge(node1.OutEdges, id2)
	node2.InEdges, changed2 = removeedge(node2.InEdges, id1)

	if changed1 {
		_, err = db.putNode(node1)
		if err != nil {
			return err
		}
	}

	if changed2 {
		_, err = db.putNode(node2)
		if err != nil {
			return err
		}
	}

	return nil
}

// Connect two nodes by adding edges in both directions.
func (db *GraphDb) Connect(id1, id2 NodeId) error {
	node1, err := db.Get(id1)
	if err != nil {
		return err
	}

	node2, err := db.Get(id2)
	if err != nil {
		return err
	}

	var changed1, changed2, changed3, changed4 bool

	node1.OutEdges, changed1 = addedge(node1.OutEdges, id2)
	node1.InEdges, changed2 = addedge(node1.InEdges, id2)
	node2.OutEdges, changed3 = addedge(node2.OutEdges, id1)
	node2.InEdges, changed4 = addedge(node2.InEdges, id1)

	if changed1 || changed2 {
		_, err = db.putNode(node1)
		if err != nil {
			return err
		}
	}

	if changed3 || changed4 {
		_, err = db.putNode(node2)
		if err != nil {
			return err
		}
	}

	return nil
}

// Connect two nodes by adding edges in both directions.
func (db *GraphDb) Disconnect(id1, id2 NodeId) error {
	node1, err := db.Get(id1)
	if err != nil {
		return err
	}

	node2, err := db.Get(id2)
	if err != nil {
		return err
	}

	var changed1, changed2, changed3, changed4 bool

	node1.OutEdges, changed1 = removeedge(node1.OutEdges, id2)
	node1.InEdges, changed2 = removeedge(node1.InEdges, id2)
	node2.OutEdges, changed3 = removeedge(node2.OutEdges, id1)
	node2.InEdges, changed4 = removeedge(node2.InEdges, id1)

	if changed1 || changed2 {
		_, err = db.putNode(node1)
		if err != nil {
			return err
		}
	}

	if changed3 || changed4 {
		_, err = db.putNode(node2)
		if err != nil {
			return err
		}
	}

	return nil
}
