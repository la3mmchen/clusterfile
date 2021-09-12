package app

import (
	"github.com/la3mmchen/clusterfile/internal/types"
	"github.com/urfave/cli/v2"
)

// CreateApp builds the cli app by enriching the
// urface/cli app struct with our params, flags, and commands.
// returns a pointer to a cli.App struct
func CreateApp(cfg *types.Configuration) *cli.App {
	cliFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "clusterfile",
			Value:       cfg.ClusterfileLocation,
			Usage:       "Clusterfile to parse.",
			Destination: &cfg.ClusterfileLocation,
		},
		&cli.StringFlag{
			Name:        "env",
			Value:       "",
			Usage:       "Select an environment to operate instead of all.",
			Destination: &cfg.EnvSelection,
		},
		&cli.StringFlag{
			Name:        "helmfile-binary",
			Value:       "/usr/local/bin/helmfile",
			Usage:       "Helmfile executable to use.",
			Destination: &cfg.HelmfileExecutable,
		},
		&cli.StringFlag{
			Name:        "output-dir",
			Value:       ".rendered",
			Usage:       "Output-Dir to write to.",
			Destination: &cfg.OutputDir,
		},
		&cli.StringFlag{
			Name:        "kube-context",
			Value:       cfg.OverwrittenKubeContext,
			Usage:       "Overwrite kubernetes context. (may be useful for ci/cd.)",
			Destination: &cfg.OverwrittenKubeContext,
		},
		&cli.BoolFlag{
			Name:        "ignore",
			Value:       false,
			Destination: &cfg.Ignore,
			DefaultText: "default: false, if true ignore missing helmfile.",
			Usage:       "Ignore missing env specific helmfile.",
		},
	}

	cliFlags = append(cliFlags, cfg.AdditionalFlags...)

	app := cli.App{
		Name:    cfg.AppName,
		Usage:   cfg.AppUsage,
		Version: cfg.AppVersion,
		Commands: []*cli.Command{
			Build(cfg),
			Destroy(cfg),
			Diff(cfg),
			Dump(cfg),
			HelmTest(cfg),
			Lint(cfg),
			List(cfg),
			Preflight(cfg),
			Status(cfg),
			Sync(cfg),
			Template(cfg),
		},
		Flags: cliFlags,
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	return &app
}
