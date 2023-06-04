package daemon

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteCommand(t *testing.T) {
	ipfsPath := os.Getenv("IPFS_PATH")
	if ipfsPath == "" {
		t.Skip("IPFS_PATH environment variable is not set. Skipping the test.")
	}

	output := bytes.Buffer{}

	cmd := exec.Command("ipfs", "id")

	// Set the command's stdout and stderr to the buffer
	cmd.Stdout = &output
	cmd.Stderr = &output

	ExecuteCommand()

	expectedOutput := "ipfs id\n"
	actualOutput := strings.TrimSpace(output.String())
	assert.Equal(t, expectedOutput, actualOutput)
}
