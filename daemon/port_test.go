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

	parsedPort, parseErr := strconv.Atoi(port)
	if parseErr != nil {
		t.Fatalf("returned port is not a valid integer: %v", parseErr)
	}

	if parsedPort < 8080 || parsedPort > 8999 {
		t.Fatalf("returned port is not within the expected range")
	}

	ln, listenErr := net.Listen("tcp", ":"+port)
	if listenErr != nil {
		t.Fatalf("failed to listen on the returned port: %v", listenErr)
	}
	defer ln.Close()

	if ln.Addr().(*net.TCPAddr).Port != parsedPort {
		t.Fatalf("the port returned by Find() is not open")
	}
}

func TestAssign(t *testing.T) {
	got, err := Assign()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// want := 
}

func TestAssignToGateway(t *testing.T) {

}

func TestAssignToAPI(t *testing.T) {
	
}