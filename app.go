package main

import "os"

// Holder for operating system commands
type App interface {
	Exit(code int)
	Args() []string
}

// The default implementation
type CliApp struct{}

// Runs os.Exit(int)
func (a *CliApp) Exit(code int) {
	os.Exit(code)
}

func (a *CliApp) Args() []string {
	return os.Args
}
