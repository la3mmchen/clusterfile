package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/la3mmchen/clusterfile/internal/commands"
	"github.com/la3mmchen/clusterfile/internal/helpers"
	"github.com/urfave/cli/v2"
)

func bootstrapTestApp() *cli.App {
	// construct an app for testing purposes
	cfg := helpers.GetTestCfg()

	return commands.CreateApp(&cfg)
}

func TestApp(t *testing.T) {

	app := bootstrapTestApp()

	args := os.Args[0:1]
	args = append(args, "")

	if err := app.Run(args); err != nil {
		fmt.Printf("Error: %v \n", err)
		os.Exit(1)
	}
}

func TestSubcmdBuild(t *testing.T) {
	app := bootstrapTestApp()

	args := os.Args[0:1]
	args = append(args, "build")

	if err := app.Run(args); err != nil {
		fmt.Printf("Error: %v \n", err)
		os.Exit(1)
	}
}

func TestSubcmdDump(t *testing.T) {
	app := bootstrapTestApp()

	args := os.Args[0:1]
	args = append(args, "dump")

	if err := app.Run(args); err != nil {
		fmt.Printf("Error: %v \n", err)
		os.Exit(1)
	}
}

func TestSubcmdPreflight(t *testing.T) {
	app := bootstrapTestApp()

	args := os.Args[0:1]
	args = append(args, "preflight", "--offline")

	if err := app.Run(args); err != nil {
		fmt.Printf("Error: %v \n", err)
		os.Exit(1)
	}
}
