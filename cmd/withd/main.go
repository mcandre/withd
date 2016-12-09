package main

import (
	"os"
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
	arguments, _ := docopt.Parse(Usage, nil, true, withd.Version, false)

	directory, _ := arguments["<chdir>"].(string)
	commandString, _ := arguments["<command>"].(string)
	args, _ := arguments["<args>"].([]string)

	command := exec.Command(commandString, args...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	command.Dir = directory
	command.Run()
}
