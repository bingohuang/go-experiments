package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	Start()
}

// Start starts a nginx (master process) and waits. If the process ends
// we need to kill the controller process and return the reason.
func Start() {
	fmt.Println("Starting cseproxy process...")
	logdir := "/data/test/exec"
	outlogfile, err := os.OpenFile(logdir+"/start.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0640)
	if err != nil {
		fmt.Errorf("create logs file failed, %v", err)
		panic("create logs file failed.")
	}

	cmd := exec.Command("sh", "-c", "/data/test/exec/start.sh")
	cmd.Stdout = outlogfile
	cmd.Stderr = outlogfile
	if err := cmd.Start(); err != nil {
		fmt.Errorf("nginx error: %v", err)
		panic("nginx try to restart failed.")
	}

	if err := cmd.Wait(); err != nil {
		fmt.Errorf("nginx error: %v", err)
		//panic("nginx wait failed.")
	}
}
