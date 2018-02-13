package graphdb

type WeaveType struct {
  Id NodeId
  Node *Node
}

// Connect all nodes on the left to all nodes on the right bidirectionally.
//
// This is an efficientcy as it only writes each node once.
//
func (db *GraphDb) Weave(left, right []WeaveType) {
  leftChanged := make([]bool, len(left))
  rightChanged := make([]bool, len(right))

  for lidx, leftWeave := range left {
    for ridx, rightWeave := range right {
      var changed1, changed2, changed3, changed4 bool

      leftWeave.Node.OutEdges, changed1 = addedge(leftWeave.Node.OutEdges, rightWeave.Id)
    	leftWeave.Node.InEdges, changed2 = addedge(leftWeave.Node.InEdges, rightWeave.Id)
    	rightWeave.Node.OutEdges, changed3 = addedge(rightWeave.Node.OutEdges, leftWeave.Id)
    	rightWeave.Node.InEdges, changed4 = addedge(rightWeave.Node.InEdges, leftWeave.Id)

      leftChanged[lidx] = changed1 || changed2
      rightChanged[ridx] = changed3 || changed4
    }
  }

  for i, changed := range leftChanged {
    if changed {
      node := left[i].Node
      db.Put(node.Name, node.Data, node.InEdges, node.OutEdges)
    }
  }

  for i, changed := range rightChanged {
    if changed {
      node := right[i].Node
      db.Put(node.Name, node.Data, node.InEdges, node.OutEdges)
    }
  }
}

// Connect all nodes on the left to all nodes on the right unidirectionally.
//
// This is an efficientcy as it only writes each node once.
//
func (db *GraphDb) UniWeave(left, right []WeaveType) {
  leftChanged := make([]bool, len(left))
  rightChanged := make([]bool, len(right))

  for lidx, leftWeave := range left {
    for ridx, rightWeave := range right {
      var changed1, changed2 bool

      leftWeave.Node.OutEdges, changed1 = addedge(leftWeave.Node.OutEdges, rightWeave.Id)
    	rightWeave.Node.InEdges, changed2 = addedge(rightWeave.Node.InEdges, leftWeave.Id)

      leftChanged[lidx] = changed1
      rightChanged[ridx] = changed2
    }
  }

  for i, changed := range leftChanged {
    if changed {
      node := left[i].Node
      db.Put(node.Name, node.Data, node.InEdges, node.OutEdges)
    }
  }

  for i, changed := range rightChanged {
    if changed {
      node := right[i].Node
      db.Put(node.Name, node.Data, node.InEdges, node.OutEdges)
    }
  }
}
