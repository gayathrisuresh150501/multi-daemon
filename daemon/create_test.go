package daemon

// import (
// 	"os"
// 	"testing"
// )

// var DirPath, _ = os.UserHomeDir()
// var Dname = `.ipfs`

// var Dir = DirPath + `\` + Dname


// func TestGetHomeDir(t *testing.T) {
// 	got, _:= GetHomeDir()
// 	want, _ := os.UserHomeDir()

// 	if got != want {
// 		t.Errorf("Could not fetch Home Directory Path.\ngot = %v, want = %v", got, want)
// 	}
// }

// func TestDirExistance(t *testing.T) {
// 	os.Chdir(DirPath)
// 	got, _ := CheckDirExistance(Dir)

// 	want := true

// 	if got != want {
// 		t.Errorf("Directory does not exist.\ngot = %v, want = %v", got, want)
// 	}

// }

// func TestCreateRootInst(t *testing.T) {
// 	os.Chdir(DirPath)
// 	t.Run("Root instance exists", func(t *testing.T) {
// 		got := CreateRootInst()
// 		var want error = nil

// 		if got != nil {
// 			t.Errorf("Root instance does not exist.\ngot = %v, want = %v", got, want)
// 		}
// 	})

// 	t.Run("Root instance does not exist, create root instance", func(t *testing.T) {
// 		got := CreateRootInst()

// 	})

// }

// func TestCreateNewInstances(t *testing.T) {
// 	os.Chdir(DirPath)
// 	got := CreateNewInstances(InstCount)
	
// }