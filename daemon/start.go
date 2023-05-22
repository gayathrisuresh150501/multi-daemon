package daemon

import "os/exec"

func Start() error{
	for _, inst := range InstWithPaths{
		SetPath(InstWithPaths[inst])
		exec.Command("ipfs", "daemon")
	}

	return ErrUnableToRunDaemon
}