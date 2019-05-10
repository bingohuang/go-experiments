package main

import (
	"encoding/json"
	"fmt"
)

//type Person struct {
//	Name   string
//	Age    int
//	Emails []string
//}

type Person struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Emails []string `json:"emails"`
}

func main() {
	bingo := Person{
		Name:   "Bingo Huang",
		Age:    30,
		Emails: []string{"go@bingohuang.com", "me@bingohuang.com"},
	}

	json_bytes, err := json.Marshal(bingo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s",json_bytes)
}
