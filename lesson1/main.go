package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Message ...
type Message struct {
	Body string
}

func main() {

	input := bufio.NewReader(os.Stdin)
	line, _, _ := input.ReadLine()

	output := bufio.NewWriter(os.Stdout)
	n, _ := output.Write(line)
	text := string(line[:n])

	m := Message{text}
	b, _ := json.Marshal(m)
	fmt.Println(string(b[:len(b)]))

}
