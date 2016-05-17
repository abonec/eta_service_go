package main

import "testing"

func TestNewDbQuery(t *testing.T) {
	query := NewDbQuery()
	if query.indexType != "cab" {
		t.Error("Index type should be a cab")
	}
	if query.etaModifier != 1.5 {
		t.Error("etaModifier should be a 1.5")
	}
}

