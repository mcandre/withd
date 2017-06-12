// Package main provides a withd executable for evaluating commands in a specific current working directory.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/mcandre/withd"
)

var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

// main is the entrypoint for this application.
func main() {
	flag.Parse()

	switch {
	case *flagHelp:
		flag.PrintDefaults()
		os.Exit(1)
	case *flagVersion:
		fmt.Println(withd.Version)
		os.Exit(0)
	}

	args := flag.Args()

	if len(args) < 2 {
		log.Panic("A directory and a command must be supplied, following any named flags")
	}

	directory, command, commandOptions := args[0], args[1], args[2:]

	commandPlan := exec.Command(command, commandOptions...)
	commandPlan.Stderr = os.Stderr
	commandPlan.Stdout = os.Stdout
	commandPlan.Dir = directory

	err := commandPlan.Run()

	if err != nil {
		log.Panic(err)
	}
}
