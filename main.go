package main

import (
	"fmt"
	"log"
	"package/daemon"
)

func main() {
	fmt.Print("Please enter the number of daemons to be spawned: ")
	nodeCount, err := daemon.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = daemon.CreateNewInstances(nodeCount) //Init() Set() Port
	if err != nil {
		log.Fatal(err)
	}

	daemon.Start()
	

	response := daemon.GetResponse()
	if response {
		daemon.Shutdown()
	}

}
