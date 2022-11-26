package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("myapp.log")
	// Nil -> Empty pointer values. Similar to null
	if err != nil {
		log.Fatal(err)
	}
	// Defer -> Close this when our function is done with the resource
	defer f.Close()

	// Read contents of file for purpose of why we even made the analyzer
	// bufio (buffered input output package)
	// bufio.NewReader -> Input stream reader. Pretty self explanatory. Reads files, network connection, input string, etc.
	r := bufio.NewReader(f)
	// Use Go's infinite for loop
	for {
		// Parse on newline char
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, "ERROR") {
			fmt.Println(s)
		}
	}
}
