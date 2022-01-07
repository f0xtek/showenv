package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	// create a new reader to read keyboard input from the stdin stream
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What is your name?")
	// read the stdin stream until a newline character is received
	text, _ := reader.ReadString('\n')

	fmt.Printf("Hello %v", text) // the newline character scanned by reader is kept and added to the output!
	fmt.Printf("We are using Go %v in %v!\n", runtime.Version(), runtime.GOOS)
}
