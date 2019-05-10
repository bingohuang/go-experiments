package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string      `json:"name,omitempty"`
	Age    int         `json:"-"`
	Emails []string    `json:"emails,omitempty"`
}

func main() {
	bingo := Person{
		Name: "Bingo Huang",
		Age:  30,
	}

	json_bytes, _ := json.Marshal(bingo)
	fmt.Printf("%s", json_bytes)
}