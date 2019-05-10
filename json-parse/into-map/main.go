package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	json_bytes := []byte(`
		{
			"Name":"Bingo Huang",
			"Age":30,
			"Emails":["go@bingohuang.com","me@bingohuang.com"]
		}
	`)

	var bingo map[string]interface{}
	err := json.Unmarshal(json_bytes, &bingo)
	if err != nil {
		panic(err)
	}

	fmt.Println(bingo["Name"], bingo["Age"], bingo["Emails"], bingo["Score"])

}