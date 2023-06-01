package daemon

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteCommand() {
	// defer wg.Done()

	// var cmd *exec.Cmd
	cmd := exec.Command("ipfs", "daemon")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		panic(err)
		// log.Fatal(err)
	}

	execPath := os.Getenv("IPFS_PATH")

	fmt.Printf("Command: %s\nOutput: %s\n", "ipfs daemon on", execPath)
}
