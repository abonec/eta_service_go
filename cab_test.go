package main

import "testing"

func TestNewCabFromCoordinates(t *testing.T) {
	cab := NewCabFromCoordinates("1.5, 2.6", true)
	if cab.Location.Lat != 1.5 || cab.Location.Lon != 2.6 {
		t.Errorf("Cab coordinates should be {1.5, 2.6} but was {%.2f, %.2f}", cab.Location.Lat, cab.Location.Lon)
	}

	if !cab.Vacant {
		t.Error("Cab should be vacant")
	}
}

func TestToCorrectUnmarshal(t *testing.T) {
	cab := NewCabFromJson([]byte(`{"vacant": true, "location": {"lat": 1.5, "lon": 2.6}}`))
	if !cab.Vacant {
		t.Error("Cab should be vacant")
	}

	if cab.Location.Lat != 1.5 || cab.Location.Lon != 2.6 {
		t.Errorf("Cab coordinates should be {1.5, 2.6} but was {%.2f, %.2f}", cab.Location.Lat, cab.Location.Lon)
	}
}
