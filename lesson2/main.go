package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

// Message ...
type Message struct {
	Body string
}

var addr = flag.String("address", "localhost:5000", "address to connect to")

func init() {
	flag.Parse()
}

func main() {

	conn, err := net.Dial("tcp", *addr)

	if err != nil {

		log.Fatalln("Could not connect to ", *addr, " because of: ", err)
	}
	for {

		input := bufio.NewReader(os.Stdin)
		fmt.Print("What message to send: ")
		line, _, _ := input.ReadLine()

		output := bufio.NewWriter(os.Stdout)
		n, _ := output.Write(line)
		text := string(line[:n])
		// something

		m := Message{text}
		b, _ := json.Marshal(m)
		fmt.Println(string(b[:len(b)]))
		fmt.Println("hello world")
		fmt.Fprintf(conn, string(b[:len(b)]))

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)

	}
}
