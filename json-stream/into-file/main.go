package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
}

func main() {
	bingo := Person{
		Name:   "Bingo Huang",
		Age:    30,
		Emails: []string{"go@bingohuang.com","me@bingohuang.com"},
	}

	fileWriter, _ := os.Create("output.json")
	json.NewEncoder(fileWriter).Encode(bingo)
}