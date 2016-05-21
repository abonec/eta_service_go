package main

import (
	"testing"
)

func TestReadCabs(t *testing.T) {
	cabs := ReadCabs(cab_fixtures_file_path, cab_fixtures_size)
	if len(cabs) != cab_fixtures_size {
		t.Error("ReadCabs should import only 1000 examples")
	}

	first_cab := cabs[0]
	last_cab := cabs[len(cabs)-1]

	if first_cab.Location.Lat == 0 || first_cab.Location.Lon == 0 || !first_cab.Vacant ||
		last_cab.Location.Lat == 0 || last_cab.Location.Lon == 0 || !last_cab.Vacant {
		t.Error("Nil or wrong vacant exported in first or last cab")
	}
}
