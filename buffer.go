package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.WriteString("Hello")
	buffer.WriteString(" World")
	fmt.Printf("buffer: %v\n", buffer.String())
}
