package daemon

import (
	"os"
	"os/exec"
)

func ExecCmd(ipfs, ipfsCmd string) error{
	cmd := exec.Command(ipfs, ipfsCmd)

	// Set command output to os.Stdout and os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return ErrUnableToRunDaemon
	}

	// err = cmd.Wait()
	// if err != nil {
	// 	return err
	// }

	return nil
}
