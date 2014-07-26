/*
An example CLI application.

Installation:
	go get github.com/dickeyxxx/programming-in-go
*/
package main

import (
	"fmt"
	"io"
	"os"
)

// Default config dependencies
var (
	stdout io.Writer = os.Stdout
	stderr io.Writer = os.Stderr
	app    App       = &CliApp{}
)

// The available commands
var Commands = []string{"run"}

// Main entry point
func main() {
	if len(app.Args()) > 1 {
		switch app.Args()[1] {
		case "run":
			Run()
		default:
			printUsage()
			app.Exit(1)
		}
	} else {
		printUsage()
		app.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintln(stderr, "USAGE:", app.Args()[0], "[command]")
	fmt.Fprintln(stderr, "Where [command] is one of:")
	for _, command := range Commands {
		fmt.Fprintln(stderr, command)
	}
}

// Main example command
func Run() {
	fmt.Fprintln(stdout, "running!")
}
