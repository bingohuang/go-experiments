package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Emails []string `json:"emails"`
}

func main() {
	json_bytes := []byte(`
		{
			"Name":"Bingo Huang",
			"Age":30,
			"Emails":["go@bingohuang.com","me@bingohuang.com"]
		}
	`)

	bingo := Person{}
	err := json.Unmarshal(json_bytes, &bingo)
	if err != nil {
		panic(err)
	}

	fmt.Println(bingo.Name, bingo.Age, bingo.Emails)

}