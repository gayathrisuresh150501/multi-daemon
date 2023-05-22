package daemon

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	RootConfigDir = ".ipfs"
)

var HomeDir, _ = GetHomeDir()
var Path = HomeDir + "/" + RootConfigDir

func GetHomeDir() (string, error) {
	return os.UserHomeDir()
}

func CheckDirExistance(rootPath string) (bool, error) {
	if _, err := os.Stat(rootPath); err != nil {
		if err.Error() == "no such file or directory" { //resolve this error in a better
			return false, nil
		}

		return false, err
	}

	return true, nil
}

var InstWithPaths map[string]string

func CreateRootInst() error {
	exists, _ := CheckDirExistance(Path)
	if !exists {
		exec.Command("ipfs", "init")
	}
	InstWithPaths[RootConfigDir] = Path

	return ErrCouldNotCreateRootInst
}
func CreateNewInstances(instanceCount int) error {
	os.Chdir(HomeDir)

	for i := 1; i < instanceCount; i++ {
		instName := RootConfigDir + fmt.Sprintf("%d",i)

		exists, err := CheckDirExistance(instName)
		if err == nil && !exists {
			os.Mkdir(instName, 0755)
			instPath := HomeDir + instName
			InstWithPaths[instName] = instPath
		} else {
			return ErrDirNotFound
		}

	}
	err := Init(InstWithPaths)
	if err != nil {
		return ErrCouldNotCreateInst
	}
	//check return val
	return nil
}
