package main

import (
	"fmt"
	_ "golang.org/x/net"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}

	fmt.Println(interfaces)
}
