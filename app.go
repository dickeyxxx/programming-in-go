package main

import "os"

type App interface {
	Exit(code int)
	Args() []string
}

type CliApp struct{}

func (a *CliApp) Exit(code int) {
	os.Exit(code)
}

func (a *CliApp) Args() []string {
	return os.Args
}
