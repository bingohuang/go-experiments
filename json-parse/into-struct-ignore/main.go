package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Emails  []string `json:"emails"`
	Address string
}

func main() {

	json_bytes := []byte(`
		{
			"Name":"Bingo Huang",
			"Age":30
		}
	`)

	var bingo Person
	err := json.Unmarshal(json_bytes, &bingo)
	if err != nil {
		panic(err)
	}

	fmt.Println(bingo.Emails)
	fmt.Println(bingo.Address)

}