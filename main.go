package main

import (
	"fmt"
	"net"
	"os"
)

const numOfArgs = 5

func main() {
	if len(os.Args) != numOfArgs {
		os.Exit(-1)
	}
	fmt.Println(net.JoinHostPort(os.Args[1], os.Args[2]))
}
