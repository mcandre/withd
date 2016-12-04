package main

import (
	"fmt"
	"os/exec"

	"github.com/docopt/docopt-go"
	"github.com/mcandre/withd"
)

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

func main() {
	arguments, err := docopt.Parse(Usage, nil, true, withd.Version, false)

	if err != nil {
		panic(Usage)
	}

	directory, _ := arguments["<chdir>"].(string)

	if directory == "" {
		panic(Usage)
	}

	commandString, _ := arguments["<command>"].(string)

	if commandString == "" {
		panic(Usage)
	}

	args, _ := arguments["<args>"].([]string)

	command := exec.Command(commandString, args...)
	command.Dir = directory

	output, err := command.Output()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", output)
}
