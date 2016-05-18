package main

import (
	"testing"
)

func TestReadCabs(t *testing.T) {
	cabs := ReadCabs("fixtures/cabs.txt", 1000)
	if len(cabs) != 1000 {
		t.Error("ReadCabs should import only 1000 examples")
	}

	first_cab := cabs[0]
	last_cab := cabs[len(cabs)-1]

	if first_cab.Location.Lat == 0 || first_cab.Location.Lon == 0 || !first_cab.Vacant ||
		last_cab.Location.Lat == 0 || last_cab.Location.Lon == 0 || !last_cab.Vacant {
		t.Error("Nil or wrong vacant exported in first or last cab")
	}
}