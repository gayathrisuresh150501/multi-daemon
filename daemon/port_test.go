package daemon

import (
	"net"
	"strconv"
	"testing"
)

func TestFind(t *testing.T) {
	port, err := Find()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Checks if the returned port is valid
	parsedPort, parseErr := strconv.Atoi(port)
	if parseErr != nil {
		t.Fatalf("returned port is not a valid integer: %v", parseErr)
	}

	// Checks if the returned port is within the expected range
	if parsedPort < 80000 || parsedPort > 90000 {
		t.Fatalf("returned port is not within the expected range")
	}

	// Attempts to listen on the returned port
	ln, listenErr := net.Listen("tcp", ":"+port)
	if listenErr != nil {
		t.Fatalf("failed to listen on the returned port: %v", listenErr)
	}
	defer ln.Close()

	// Checks if the port is actually open
	if ln.Addr().(*net.TCPAddr).Port != parsedPort {
		t.Fatalf("the port returned by Find() is not open")
	}
}
