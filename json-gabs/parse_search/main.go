package main

import (
	"fmt"
	"github.com/Jeffail/gabs"
)

func main() {
	jsonParsed, _ := gabs.ParseJSON([]byte(`{
	"outter":{
		"inner":{
			"value1":10,
			"value2":22
		},
		"alsoInner":{
			"value1":20
		}
	}
	}`))

	var value float64
	var ok bool

	value, ok = jsonParsed.Path("outter.inner.value1").Data().(float64)
	// value == 10.0, ok == true
	fmt.Println("value:", value)
	fmt.Println("ok:", ok)

	value, ok = jsonParsed.Search("outter", "inner", "value1").Data().(float64)
	// value == 10.0, ok == true
	fmt.Println("value:", value)
	fmt.Println("ok:", ok)

	value, ok = jsonParsed.Path("does.not.exist").Data().(float64)
	// value == 0.0, ok == false
	fmt.Println("value:", value)
	fmt.Println("ok:", ok)

	exists := jsonParsed.Exists("outter", "inner", "value1")
	// exists == true
	fmt.Println("exists:", exists)

	exists = jsonParsed.Exists("does", "not", "exist")
	// exists == false
	fmt.Println("exists:", exists)

	exists = jsonParsed.ExistsP("does.not.exist")
	// exists == false
	fmt.Println("exists:", exists)
}
