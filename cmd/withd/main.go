// Package main provides a withd executable for evaluating commands in a specific current working directory.
package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/docopt/docopt-go"
	"github.com/mcandre/withd"
)

// Usage is a docopt-formatted specification of this application's command line interface.
const Usage = `Usage:
  withd <chdir> <command> [<args>]...
  withd -h
  withd -v

  Arguments:
    <chdir>       The directory to treat as CWD
    <command>     The command to run
    <args>        Any flags to pass to the command
  Options:
    -h --help     Show usage information
    -v --version  Show version information`

// main is the entrypoint for this application.
func main() {
	arguments, _ := docopt.Parse(Usage, nil, true, withd.Version, false)

	directory, _ := arguments["<chdir>"].(string)
	commandString, _ := arguments["<command>"].(string)
	args, _ := arguments["<args>"].([]string)

	command := exec.Command(commandString, args...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	command.Dir = directory
	err := command.Run()

	if err != nil {
		log.Panic(err)
	}
}
