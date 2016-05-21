package main

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
type Cab struct {
	Vacant   bool     `json:"vacant"`
	Location Location `json:"location"`
}

func NewCabFromCoordinates(coordinates string, vacant bool) *Cab {
	split := strings.Split(coordinates, ",")
	lat, err := strconv.ParseFloat(strings.TrimSpace(split[0]), 64)
	HandleError(err)
	lon, err := strconv.ParseFloat(strings.TrimSpace(split[1]), 64)
	HandleError(err)
	cab := &Cab{
		Vacant: vacant,
		Location: Location{
			Lat: lat,
			Lon: lon,
		},
	}
	return cab
}

func NewCabFromJson(json_bytes []byte) *Cab {
	cab := &Cab{}
	err := json.Unmarshal(json_bytes, cab)
	HandleError(err)
	return cab
}

func (cab *Cab) ToJson() []byte {
	result, err := json.Marshal(cab)
	HandleError(err)
	return result
}
