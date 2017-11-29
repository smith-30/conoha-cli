package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mitchellh/cli"
	"github.com/smith-30/conoha-cli/subcmd"
)

var (
	revision = "unknown"
)

func main() {
	LoadEnv()

	c := cli.NewCLI("conoha-cli", revision)

	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"boot": func() (cli.Command, error) {
			return &subcmd.Boot{}, nil
		},
		"halt": func() (cli.Command, error) {
			return &subcmd.Halt{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
