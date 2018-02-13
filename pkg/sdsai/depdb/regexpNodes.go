package depdb

import (
	"bytes"
	"encoding/base64"
	"github.com/basking2/sdsai-graphdb-golang/pkg/sdsai/graphdb"
	"regexp"
)

var regexp_key_prefix = "regexp_"

func re2key(re *regexp.Regexp) string {
	buf := bytes.Buffer{}
	enc := base64.NewEncoder(base64.RawStdEncoding, &buf)
	enc.Write([]byte(re.String()))
	enc.Close()
	key := graphdb.ScrambleKey(buf.String())

	return regexp_key_prefix + key
}

func (db *DepDb) FindAllRegexp() ([]graphdb.NodeId, error) {
	ids, err := db.graphdb.FindIds(regexp_key_prefix)
	if err != nil {
		return make([]graphdb.NodeId, 0), err
	}

	deps := make([]graphdb.NodeId, len(ids))

	for i, id := range ids {
		deps[i] = id
	}

	return deps, nil
}

// Add or update a mapping from a regular expression node to a basic node.
func (db *DepDbBuilder) AddRegexpMapping(re *regexp.Regexp, basicId string) error {

	db.logger.Infof("Mapping %s to %s.", re.String(), basicId)

	basicDepData := DepData{}

	basicDepId, _, err := db.graphdb.FindOneData(basic2key(basicId), &basicDepData)
	if err == graphdb.NotFoundError {
		basicDepData.Name = basicId
		_, err = db.graphdb.Put(basic2key(basicId), &basicDepData, nil, nil)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	reDepData := DepData{}
	reDepId, _, err := db.graphdb.FindOneData(re2key(re), &reDepData)
	if err == graphdb.NotFoundError {
		reDepData.Name = re.String()
		reDepData.Regexp = re.String()
		_, err = db.graphdb.Put(re2key(re), &reDepData, nil, nil)
		if err != nil {
			return err
		}
	}

	err = db.graphdb.AddEdge(reDepId, basicDepId)
	if err != nil {
		return err
	}

	db.regexps = append(db.regexps, re)

	return nil
}
