package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	lsn, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln("Could not establish an incoming connection: ", err)
	}
	for {
		c, err := lsn.Accept()

		if err != nil {
			log.Fatalln("Accepting the connection failed : ", err)

		}

		fmt.Fprintln(c, "Welcome to the server")

		io.Copy(c, c)

	}
}
