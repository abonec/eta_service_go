package main

import "log"

func (finder *DbQuery) PutMappings() *DbQuery {
	mapping := `{
		"cab":{
			"properties":{
				"location":{
					"type":"geo_point"
				},
				"vacant":{
					"type":"boolean"
				}
			}
		}
	}`
	_, err := finder.client.PutMapping().Index(finder.index).Type(finder.indexType).BodyString(mapping).Do()
	HandleError(err)
	return finder
}

func (finder *DbQuery) CreateIndex() *DbQuery {
	_, err := finder.client.CreateIndex(finder.index).Do()
	HandleError(err)
	log.Printf("Index %s created", finder.index)
	return finder
}

func (finder *DbQuery) DestroyIndex() *DbQuery {
	_, err := finder.client.DeleteIndex(finder.index).Do()
	HandleError(err)
	log.Printf("Index %s destroyed", finder.index)
	return finder
}

func (finder *DbQuery) DestroyIndexIfExists() *DbQuery {
	if finder.IndexExists() {
		finder.DestroyIndex()
	}
	return finder
}

func (finder *DbQuery) RecreateIndex() *DbQuery {
	finder.DestroyIndexIfExists()
	finder.CreateIndex()
	finder.PutMappings()

	return finder
}
func (finder *DbQuery) IndexExists() bool {
	exist, err := finder.client.IndexExists(finder.index).Do()
	if err != nil {
		return false
	} else {
		return exist
	}
}

func (finder *DbQuery) Migrate(fixtures_size int) *DbQuery {
	finder.RecreateIndex()
	cabs := ReadCabs(cab_fixtures_file_path, fixtures_size)
	finder.BulkIndex(cabs, false, false)
	log.Printf("Cabs are imported")
	return finder
}


