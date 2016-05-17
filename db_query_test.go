package main

import (
	"testing"
)

func TestMain(m *testing.M) {
	InitDatabase()
	m.Run()
}
func TestNewDbQuery(t *testing.T) {
	query := NewTestQuery()
	if query.indexType != "cab" {
		t.Error("Index type should be a cab")
	}
	if query.etaModifier != 1.5 {
		t.Error("etaModifier should be a 1.5")
	}
}

func TestDbQuery_CreateIndex(t *testing.T) {
	query := NewTestQuery()
	query.DestroyIndex()

	if query.IndexExists() {
		t.Error("test index should not exitsts")
	}
	query.CreateIndex()
	if !query.IndexExists() {
		t.Error("test index should exists")
	}
}

