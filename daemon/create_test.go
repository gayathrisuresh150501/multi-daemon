package daemon

import (
	"os"
	"testing"
)

var DirPath, _ = os.UserHomeDir()
var Dname = `.ipfs`

var Dir = DirPath + `\` + Dname


func TestGetHomeDir(t *testing.T) {
	want, err := os.UserHomeDir()
	if err != nil {
		t.Errorf("error retrieving home directory: %v", err)
	}

	got, err := GetHomeDir()
	if err != nil {
		t.Errorf("unexpected error occurred: %v", err)
	}

	if got != want {
		t.Errorf("expected home directory '%s', but got '%s'", want, got)
	}
}

func TestCheckDirExistance(t *testing.T) {
	// Create a temporary directory for testing
	dir := ".ipfs"
	err := os.Mkdir(dir, 0755)
	if err != nil {
		t.Fatalf("Failed to create .ipfs directory: %v", err)
	}
	defer os.RemoveAll(dir)

	// Test case 1: Existing directory
	exists, err := CheckDirExistance(dir)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v", err)
	}
	if !exists {
		t.Errorf("Expected directory to exist, but it does not")
	}

	// Test case 2: Non-existing directory
	nonExistingDir := ".ipfs10"
	exists, err = CheckDirExistance(nonExistingDir)
	if err != nil {
		t.Errorf("Unexpected error occurred: %v", err)
	}
	if exists {
		t.Errorf("Expected directory to not exist, but it does")
	}
}



func TestCreateRootInst(t *testing.T) {
	os.Chdir(DirPath)
	t.Run("Root instance exists", func(t *testing.T) {
		got := CreateRootInst()
		var want error = nil

		if got != nil {
			t.Errorf("Root instance does not exist.\ngot = %v, want = %v", got, want)
		}
	})

	t.Run("Root instance does not exist, create root instance", func(t *testing.T) {
		// got := CreateRootInst()

	})

}

func TestCreateNewInstances(t *testing.T) {
	os.Chdir(DirPath)
	// got := CreateNewInstances(InstCount)
	
}