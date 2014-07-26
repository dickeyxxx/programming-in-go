package main

import (
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

var Exercises = []string{"ex1", "ex2"}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "ex1":
		// here is where real things would happen
		//Ex1([]int{9, 1, 9, )
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintln(stderr, "USAGE:", os.Args[0], "[exercise id]")
	fmt.Fprintln(stderr, "Where [exercise id] is one of:")
	for _, exercise := range Exercises {
		fmt.Fprintln(stderr, exercise)
	}
}
