package main

import (
	"strings"
	"strconv"
	"encoding/json"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
type Cab struct {
	Vacant bool `json:"vacant"`
	Location Location `json:"location"`

}

func NewCabFromCoordinates(coordinates string, vacant bool) *Cab {
	split := strings.Split(coordinates, ",")
	lat, err := strconv.ParseFloat(strings.TrimSpace(split[0]), 64)
	HandleError(err)
	lon, err := strconv.ParseFloat(strings.TrimSpace(split[1]), 64)
	HandleError(err)
	cab := &Cab {
		Vacant: vacant,
		Location: Location {
			Lat: lat,
			Lon: lon,
		},
	}
	return cab
}

func NewCabFromJson(json_string string) *Cab {
	cab := &Cab{}
	err := json.Unmarshal([]byte(json_string), cab)
	HandleError(err)
	return cab
}

