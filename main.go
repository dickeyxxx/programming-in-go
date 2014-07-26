package main

import (
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr
var app App = &CliApp{}

var Commands = []string{"run"}

func main() {
	if len(app.Args()) > 1 {
		switch app.Args()[1] {
		case "run":
			run()
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

func run() {
	fmt.Fprintln(stdout, "running!")
}
