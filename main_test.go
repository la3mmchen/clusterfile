package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/la3mmchen/clusterfile/internal/commands"
	"github.com/la3mmchen/clusterfile/internal/types"
	"github.com/urfave/cli/v2"
)

func bootstrapTestApp() *cli.App {
	// construct an app for testing purposes
	var cfg = types.Configuration{
		AppName:         "test0r",
		AppVersion:      "1234",
		AppUsage:        "Control the content of multiple k8s cluster via helmfile.",
		SkipFlagParsing: true,
		AdditionalFlags: []cli.Flag{
			&cli.StringFlag{
				Name: "test.testlogfile",
			},
			&cli.StringFlag{
				Name: "test.paniconexit0",
			},
			&cli.StringFlag{
				Name: "test.v",
			},
		},
	}

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
