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

func Assign() {
	port := 0
	m1, err := ma.NewMultiaddr("/ip4/127.0.0.1/tcp/")

}

func AssignToGateway() {

}

func AssignToAPI() {

}
