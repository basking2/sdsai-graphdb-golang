package depdb

// Dependency database.
//
// This generalizes the use of graph db.

import (
	"github.com/basking2/sdsai-graphdb-golang/pkg/sdsai/graphdb"
	"github.com/op/go-logging"
	"github.com/graymeta/stow"
	"github.com/graymeta/stow/local"
	"regexp"
)

type DepDbBuilder struct {
	// Store relationships between things.
	graphdb graphdb.GraphDb

	logger *logging.Logger

	// List of regular expressions to check.
	regexps []*regexp.Regexp

	// A mapping of nodes to their timeout in epoch.
	// If that timeout is reached the graph node is updated and the
	// webhooks are fired as if the basic event triggered.
	timeouts map[string]int64
}

// Database that holds dependencies.
type DepDb struct {
	// Store relationships between things.
	graphdb graphdb.GraphDb

	logger *logging.Logger

	// List of regular expressions to check.
	regexps []*regexp.Regexp

	// A mapping of nodes to their timeout in epoch.
	// If that timeout is reached the graph node is updated and the
	// webhooks are fired as if the basic event triggered.
	timeouts map[string]int64
}

// Private database structure used to record node information.
type DepData struct {
	// Human meaningful name for this thing.
	Name string `json:"name"`

	// If this node has a regexp that is non-zero, it can be triggered by a name event.
	Regexp string `json:"regexp"`

	// If arriving at this node should trigger any webhooks, this is that list.
	Webooks []string `json:"webhooks"`

	// The last time this was upated in epoc in UTC.
	Last int64 `json:"last"`

	// When this should be automatically triggered.
	Timeout int64 `json:"timeout"`
}

// Build a new dependency database.
func NewDepDbBuilder(name string) (DepDbBuilder, error) {
	config := make(stow.ConfigMap)
	config[local.ConfigKeyPath] = "/tmp"

	location, err := stow.Dial(local.Kind, config)
	if err != nil {
		return DepDbBuilder{}, err
	}

	graph, err := graphdb.NewGraphDb(name, location)
	if err != nil {
		return DepDbBuilder{}, err
	}

	db := DepDbBuilder{
		graph,
		logging.MustGetLogger("depdb"),
		make([]*regexp.Regexp, 0),
		make(map[string]int64),
	}

	return db, nil
}

func (db *DepDb) Get(id graphdb.NodeId) (*DepData, error) {
	data := DepData{}
	_, err := db.graphdb.GetData(id, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// Optimize the database.
// This is very important to avoid double-dispatches.
func (db *DepDbBuilder) Build() (DepDb, error) {
	remap := make(map[string]*regexp.Regexp)
	for _, re := range db.regexps {
		remap[re.String()] = re
	}

	regexps := make([]*regexp.Regexp, 0, len(remap))
	for _, v := range remap {
		regexps = append(regexps, v)
	}

	return DepDb{
		db.graphdb,
		db.logger,
		regexps,
		db.timeouts,
		}, nil
}
