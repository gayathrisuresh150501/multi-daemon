package daemon

import (
	"fmt"
	"math/rand"
	"net"
	"os/exec"
	"strconv"

	ma "github.com/multiformats/go-multiaddr"
)

// type Port struct {
// 	Addr, PortNo string
// }

func Find() (string, error) {
	maxPort := 9000
	minPort := 8000

	for i := 0; i < 10; i++ {
		randPort := rand.Intn(maxPort-minPort) + minPort
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", randPort))
		if err != nil {
			// fmt.Printf("Error listening on port %d: %v\n", randPort, err)
			continue
		}
		defer ln.Close()
		return strconv.Itoa(randPort), nil
	}
	return "", ErrCouldNotFindOpenPort
}

// func CheckPort(portNum int) bool {
// 	address := fmt.Sprintf("localhost:%d", portNum)
// 	conn, err := net.Dial("tcp", address)
// 	if err != nil {
// 		return false
// 	}
// 	defer conn.Close()
// 	return true
// }

// var port int

// func FindAny() (int, error) {
// 	fmt.Print("Please enter a port number: ")
// 	fmt.Scanf("%d", port)
// 	for {
// 		isOpen := CheckPort(port)
// 		if isOpen {
// 			return port, nil
// 		} else {
// 			fmt.Print("This port is not open please enter another port number: ")
// 			fmt.Scanf("%d", port)
// 		}
// 	}

// }

func Assign() (string, error) {
	m1, err := ma.NewMultiaddr("/ip4/127.0.0.1/tcp/")
	if err != nil {
		return "", err
	}
	Addr := m1.String()
	PortNo, _ := Find()

	return Addr + PortNo, nil
}

func AssignToGateway() {
	configGtwy, _ := Assign()
	// gtwyAddr := `"` + configGtwy + `"`
	fmt.Println("gateway addres :", configGtwy)
	if err := exec.Command("ipfs", "config", "Addresses.Gateway", configGtwy).Run(); err != nil {
		fmt.Println("Error on gateway addres :", configGtwy, err)
	}
}

func AssignToAPI() {
	configAPI, _ := Assign()
	// apiAddr := `"` + configAPI + `"`

	exec.Command("ipfs", "config", "Addresses.API", configAPI).Run()
}

// func PortChoice() string {
// 	fmt.Print("Assign port by\na)Generating a random port number\nb)User-typed port number based on the availability")
// 	var choice string

// 	for {
// 		fmt.Scanf("%s", &choice)

// 		switch choice {
// 		case "a", "A":
// 			Find()
// 		case "b", "B":
// 			FindAny()
// 		default:
// 			fmt.Print("Please enter a valid option")
// 		}

// 	}

// }
