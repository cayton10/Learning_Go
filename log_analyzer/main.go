package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Cmd line param region
	// package for cmd line params = flag
	path := flag.String("path", "myapp.log", "Path to log under analysis")
	level := flag.String("level", "ERROR", "Log level to search for. Options are DEBUG, INFO, ERROR, and CRITICAL")

	flag.Parse()

	f, err := os.Open(*path)
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
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
