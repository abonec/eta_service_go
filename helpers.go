package main

func HandleError(error interface{}) {
	if error != nil {
		panic(error)
	}
}
