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
var RootPath = HomeDir + `\` + RootConfigDir
var InstWithPaths = make(map[string]string)

func GetHomeDir() (string, error) {
	return os.UserHomeDir()
}

// func Stat() error {
// 	if _, err := os.Stat("hgffj"); err != nil {
// 		return fmt.Errorf("the error in Stat is %s", err.Error())
// 	}

// 	return nil
// }
func CheckDirExistance(Path string) (bool, error) {
	if _, err := os.Stat(Path); err != nil {

		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func CreateRootInst() error {
	exists, _ := CheckDirExistance(RootPath)
	if !exists {
		exec.Command("ipfs", "init").Run()
	}
	return ErrCouldNotCreateRootInst
}
func CreateNewInstances(instanceCount int) error {
	CreateRootInst()
	os.Chdir(HomeDir)
	// fmt.Print(HomeDir)
	InstWithPaths[RootConfigDir] = RootPath

	for i := 1; i < instanceCount; i++ {
		instName := RootConfigDir + fmt.Sprintf("%d", i)
		// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>", instName)

		exists, err := CheckDirExistance(instName)
		// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>",exists, err)

		if err == nil && exists {

			instPath := HomeDir + `\` + instName
			InstWithPaths[instName] = instPath
		} else {
			os.Mkdir(instName, 0755)
			instPath := HomeDir + `\` + instName
			InstWithPaths[instName] = instPath
			err := SetPath(instPath)
			if err != nil {
				return ErrUnableToSetPath
			}
			err = Init()
			if err != nil {
				return ErrCouldNotCreateInst
			}
			AssignToGateway()
			AssignToAPI()

		}

	}

	// fmt.Println(InstWithPaths)
	// err := Init(InstWithPaths)
	// if err != nil {
	// 	return ErrCouldNotCreateInst
	// }
	//check return val
	return nil
}
