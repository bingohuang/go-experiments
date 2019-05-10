package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	fileReader, _ := os.Open("bingo.json")
	var bingo map[string]interface{}
	json.NewDecoder(fileReader).Decode(&bingo)
	fmt.Println(bingo)
}