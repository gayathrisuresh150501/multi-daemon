package daemon

import (
	"fmt"
	"net"

	ma "github.com/multiformats/go-multiaddr"
)

type Port struct {
	Addr, PortNo string
}

func Find() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}

	tcpaddr, ok := listener.Addr().(*net.TCPAddr)
	if !ok {
		//log error
		fmt.Println("Invalid address")
	}

	return tcpaddr.Port, err
}

func FindAny() {

}

func Assign() (string, error) {
	port := 0
	m1, err := ma.NewMultiaddr("/ip4/127.0.0.1/tcp/")

}

func AssignToGateway() {
	Assign()

}

func AssignToAPI() {
	Assign()
}

func PortChoice() {
	fmt.Print("Assign port by\na)Generating a random port number\nb)User-typed port number based on the availability")
	var choice string
	fmt.Scanf("%s", &choice)

	switch choice {
	case "a", "A":
		Find()
	case "b", "B":
		FindAny()
	default:
		for {
			fmt.Print("Please enter a valid option")
		}
	}

}
