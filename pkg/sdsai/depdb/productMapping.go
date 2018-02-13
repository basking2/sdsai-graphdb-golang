package depdb

import (
	"github.com/basking2/sdsai-graphdb-golang/pkg/sdsai/graphdb"
)

func (db *DepDbBuilder) upsert(id string) (graphdb.NodeId, *graphdb.Node, error){

	depdata := DepData{}

	nodeid, node, err := db.graphdb.FindOneData(basic2key(id), &depdata)
	if err == graphdb.NotFoundError {
		depdata.Name = id
		nodeid, err = db.graphdb.Put(basic2key(id), &depdata, nil, nil)
		if err != nil {
			return graphdb.NodeId(""), nil, err
		}
		return nodeid, &graphdb.Node{basic2key(id), nil, nil, &depdata}, nil
	} else if err != nil {
		return graphdb.NodeId(""), nil, err
	}

	return nodeid, node, nil
}

// Create a mapping from many sources to many destinations.
func (db *DepDbBuilder) AddBasicMapping(sources, destinations []string) error {

	var upsertList = func(sources []string) []graphdb.WeaveType {
		lst := make([]graphdb.WeaveType, 0, len(sources))

		for _, prod := range sources {
			id, node, err := db.upsert(prod)
			if err == nil {
				lst = append(lst, graphdb.WeaveType{id, node})
			}
		}

		return lst
	}

	sourcesList := upsertList(sources)
	destinationsList := upsertList(destinations)

	db.logger.Infof("Importing merge for %s.", destinations[0])

	db.graphdb.UniWeave(sourcesList, destinationsList)

	return nil
}
