package main

import "reflect"

func NewTestQuery() *DbQuery {
	return NewDbQuery("cabs_test")
}

func EqualStructs(first, second interface{}) bool {
	return reflect.DeepEqual(first, second)
}
