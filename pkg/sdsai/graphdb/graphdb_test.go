package graphdb

import (
	"errors"
	"testing"
)

func TestGraphDbSimple(t *testing.T) {
	_, err := NewLocalGraphDb(t.Name(), "/tmp")
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGraphDbRw(t *testing.T) {
	db, err := NewLocalGraphDb(t.Name(), "/tmp")
	if err != nil {
		t.Error(err.Error())
	}

	id, err := db.Put("id", nil, nil, nil)
	if err != nil {
		t.Error(err.Error())
	}

	node2, err := db.Get(id)
	if err != nil {
		t.Error(err.Error())
	}

	if node2.Name != "id" {
		t.Error(errors.New("IDs don't match."))
	}
}

func TestGraphDbList(t *testing.T) {
	db, err := NewLocalGraphDb(t.Name(), "/tmp")
	if err != nil {
		t.Error(err.Error())
	}

	_, err = db.Put("id", nil, nil, nil)
	if err != nil {
		t.Error(err.Error())
	}

	nodes, err := db.Find("id")
	if err != nil {
		t.Error(err.Error())
	}

	if len(nodes) != 1 {
		t.Error("Nodes is the wrong length.")
	}

	if nodes[0].Name != "id" {
		t.Error(errors.New("IDs don't match."))
	}
}

func TestGraphDbGetData(t *testing.T) {
	type Ud struct {
		Name  string
		Value int
	}

	db, err := NewLocalGraphDb(t.Name(), "/tmp")
	if err != nil {
		t.Error(err.Error())
	}

	_, err = db.Put("id", &Ud{"hi", 7}, nil, nil)
	if err != nil {
		t.Error(err.Error())
	}

	ud := Ud{}
	ids, err := db.FindIds("id")
	if err != nil {
		t.Error(err.Error())
	}

	db.GetData(ids[0], &ud)

	if ud.Name != "hi" {
		t.Error("Name is not \"hi\".")
	}
	if ud.Value != 7 {
		t.Error("Value is not 7.")
	}
}

func TestGraphDbFindOneData(t *testing.T) {
	type Ud struct {
		Name  string
		Value int
	}

	db, err := NewLocalGraphDb(t.Name(), "/tmp")
	if err != nil {
		t.Error(err.Error())
	}

	_, err = db.Put("id", &Ud{"hi", 7}, nil, nil)
	if err != nil {
		t.Error(err.Error())
	}

	ud := Ud{}
	_, node, err := db.FindOneData("id", &ud)
	if err != nil {
		t.Error(err.Error())
	}

	if node.Name != "id" {
		t.Error("Name does not match.")
	}

	if ud.Name != "hi" {
		t.Error("Name is not \"hi\".")
	}
	if ud.Value != 7 {
		t.Error("Value is not 7.")
	}
}

func TestGraphDbFindOneDataNotFound(t *testing.T) {
	type Ud struct {
		Name  string
		Value int
	}

	db, err := NewLocalGraphDb(t.Name(), "/tmp")
	if err != nil {
		t.Error(err.Error())
	}

	_, err = db.Put("id-longer", &Ud{"hi", 7}, nil, nil)
	if err != nil {
		t.Error(err.Error())
	}

	ud := Ud{}
	_, _, err = db.FindOneData("id", &ud)
	if err == nil {
		t.Error("Expected error.")
	}
}
