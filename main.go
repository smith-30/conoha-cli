package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/smith-30/conoha-cli/helper/env"
	"github.com/smith-30/conoha-cli/subcmd"
)

var (
	revision = "unknown"
)

func main() {
	env.LoadEnv()

	c := cli.NewCLI("conoha-cli", revision)

	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		subcmd.BOOT: func() (cli.Command, error) {
			return &subcmd.Boot{}, nil
		},
		subcmd.HALT: func() (cli.Command, error) {
			return &subcmd.Halt{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
