package daemon

import (
	"os"
)

func SetPath(path string) error {
	return os.Setenv("IPFS_PATH", path)

}
