package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/la3mmchen/clusterfile/internal/commands"
	"github.com/la3mmchen/clusterfile/internal/helpers"
	"github.com/la3mmchen/clusterfile/internal/types"
)

var (
	// AppVersion Version of the app. Must be injected during the build.
	AppVersion string
	// Cfg types.Configuration
	Cfg types.Configuration
)

func main() {
	activeContext, err := helpers.GetActiveKubeContext()

	if err != nil {
		fmt.Printf("Error loading kube context: [%v] \n", err)
		os.Exit(1)
	}
	var cfg = types.Configuration{
		Debug:         "true",
		ActiveContext: activeContext,
	}

	app := &cli.App{
		Name:    "clusterfilectl",
		Usage:   "Control the content of multiple k8s cluster via helmfile.",
		Version: "AppVersion",
		Commands: []*cli.Command{commands.Status(&cfg), commands.Destroy(&cfg), commands.Sync(&cfg),
			commands.Diff(&cfg), commands.Dump(&cfg), commands.Preflight(&cfg),
			commands.Template(&cfg), commands.Lint(&cfg), commands.Build(&cfg)},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "clusterfile",
				Value:       "configs/clusterfile.yaml",
				Usage:       "Clusterfile to parse.",
				Destination: &cfg.ClusterfileLocation,
			},
			&cli.StringFlag{
				Name:        "helmfile-binary",
				Value:       "/usr/local/bin/helmfile",
				Usage:       "Executable",
				Destination: &cfg.HelmfileExecutable,
			},
			&cli.StringFlag{
				Name:        "helmfile",
				Value:       "helmfile/helmfile.yaml",
				Usage:       "Helmfile to use.",
				Destination: &cfg.Helmfile,
			},
			&cli.StringFlag{
				Name:        "output-dir",
				Value:       ".rendered",
				Usage:       "Output-Dir to write to",
				Destination: &cfg.OutputDir,
			},
			&cli.StringFlag{
				Name:        "kube-context",
				Value:       "",
				Usage:       "Overwrite kubernetes context. may be useful for ci/cd. (not implemented yet)", // TODO: implement me
				Destination: &cfg.OverwrittenKubeContext,
			},
			&cli.BoolFlag{
				Name:        "ignore",
				Value:       false,
				Destination: &cfg.Ignore,
				DefaultText: "default: false, if true ignore missing helmfile env app",
				Usage:       "Ignore missing env specific helmfile.",
			},
		},
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error: %v \n", err)
		os.Exit(1)
	}
}
