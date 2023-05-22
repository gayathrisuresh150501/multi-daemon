package daemon

import "os/exec"

func Init(instsWithPath map[string]string) error {
	for _, inst := range instsWithPath {
		err := SetPath(instsWithPath[inst])
		if err != nil {
			return ErrUnableToSetPath
		}
		exec.Command("ipfs", "init")
	}

	return nil
}


