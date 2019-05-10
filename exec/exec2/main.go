package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat", "file")
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()
	be, _ := ioutil.ReadAll(stderr)
	bo, _ := ioutil.ReadAll(stdout)

	cmd.Wait()
	fmt.Printf("%d be size \n", len(be))
	fmt.Printf("%d bo size \n", len(bo))
}
