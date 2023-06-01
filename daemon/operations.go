package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// type Daemon struct {
// 	NodeCount     int
// 	InstWithPaths map[string]string
// 	Port
// }

func Start() {
	// for _, inst := range InstWithPaths {
	// 	SetPath(InstWithPaths[inst])
	// 	exec.Command("ipfs", "daemon")
	// }

	// var wg sync.WaitGroup

	for _, path := range InstWithPaths {
		// wg.Add(1)
		if err := SetPath(path); err != nil {
			fmt.Println("SetPath Err:", err)
		}
		ExecuteCommand()
	}

	// wg.Wait()
}

func Shutdown() {
	exec.Command("taskkill /im ipfs.exe /f").Run()
}

func GetResponse() bool {
	var resp string
	fmt.Print("Would you like to shutdown the daemons? (y/n)")

	for {
		if n, err := fmt.Scanf("%v", &resp); err != nil {
			fmt.Println("scan err:", err, n)
		}
		if strings.ToLower(resp) == "y" {
			break
		}

		fmt.Print("Please enter a valid character!")
	}

	return true
}

func SetPath(path string) error {
	return os.Setenv("IPFS_PATH", path)

}

func Init() error {
	cmd := exec.Command("ipfs", "init")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	// err = cmd.Wait()
	// if err != nil {
	// 	return err
	// }

	return nil
}
