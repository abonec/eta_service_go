package main

import (
	"testing"
)

var dbQuery *DbQuery

func TestMain(m *testing.M) {
	InitDatabase()
	dbQuery = NewTestQuery()
	m.Run()
}
func TestNewDbQuery(t *testing.T) {
	if dbQuery.indexType != "cab" {
		t.Error("Index type should be a cab")
	}
	if dbQuery.etaModifier != 1.5 {
		t.Error("etaModifier should be a 1.5")
	}
}

func TestDbQuery_CreateIndex(t *testing.T) {
	dbQuery.DestroyIndexIfExists()

	if dbQuery.IndexExists() {
		t.Error("test index should not exitsts")
	}
	dbQuery.CreateIndex()
	if !dbQuery.IndexExists() {
		t.Error("test index should exists")
	}
}

func TestDbQuery_BulkIndex(t *testing.T) {
	dbQuery.DestroyIndex()
	dbQuery.CreateIndex()
	cabs := ReadCabs(cab_fixtures_file_path, cab_fixtures_size)
	dbQuery.BulkIndex(cabs, false, true)
	actualSize := dbQuery.IndexSize()
	if actualSize != cab_fixtures_size {
		t.Errorf("Cab index should be size of %d but was %d", cab_fixtures_size, actualSize)
	}
	firstCab := dbQuery.GetById(0)
	lastCab := dbQuery.GetById(len(cabs) - 1)

	if !EqualStructs(firstCab, cabs[0]) || !EqualStructs(lastCab, cabs[len(cabs)-1]) {
		t.Error("Structs getted rom elastic is wrong")
	}
}
