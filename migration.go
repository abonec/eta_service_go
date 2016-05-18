package main

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
	return finder
}

func (finder *DbQuery) DestroyIndex() *DbQuery {
	_, err := finder.client.DeleteIndex(finder.index).Do()
	HandleError(err)
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
