package daemon

import (
	"fmt"
	"os/exec"
	"strings"
)

func Shutdown() {
	exec.Command("taskkill /im ipfs.exe /f")
}

func GetResponse() bool {
	var resp string
	fmt.Print("Would you like to shutdown the daemons? (y/n)")

	for {
		fmt.Scanf("%s", resp)
		if strings.ToLower(resp) == "y" {
			break
		}

		fmt.Print("Please enter a valid character!")
	}

	return true
}
