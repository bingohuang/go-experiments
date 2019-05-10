package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	emails := []string{"go@bingohuang.com", "me@bingohuang.com"}
	json_bytes, _ := json.Marshal(emails)
	fmt.Printf("%s", json_bytes)
}