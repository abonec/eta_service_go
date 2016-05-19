package main

import "testing"

func TestDbQuery_Migrate(t *testing.T) {
	dbQuery.DestroyIndexIfExists()
	if dbQuery.IndexExists() {
		t.Error("Index should not exists before migrate")
	}
	dbQuery.Migrate(cab_fixtures_size)
	if !dbQuery.IndexExists() {
		t.Error("Index should exists after migrate")
	}
	if dbQuery.IndexSize() != cab_fixtures_size {
		t.Errorf("Index should contains %d records", cab_fixtures_size)
	}
}