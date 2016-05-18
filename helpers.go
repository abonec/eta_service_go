package main

const (
	cab_fixtures_file_path = "fixtures/cabs.txt"
	cab_fixtures_size = 1000
)
func HandleError(error interface{}) {
	if error != nil {
		panic(error)
	}
}
