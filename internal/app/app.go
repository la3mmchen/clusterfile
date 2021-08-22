package app

import (
	"github.com/la3mmchen/clusterfile/internal/types"
	"github.com/urfave/cli/v2"
)

func CreateApp(cfg *types.Configuration) *cli.App {
	cliFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "clusterfile",
			Value:       cfg.ClusterfileLocation,
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
			Value:       cfg.OverwrittenKubeContext,
			Usage:       "Overwrite kubernetes context. may be useful for ci/cd.)",
			Destination: &cfg.OverwrittenKubeContext,
		},
		&cli.BoolFlag{
			Name:        "ignore",
			Value:       false,
			Destination: &cfg.Ignore,
			DefaultText: "default: false, if true ignore missing helmfile env app",
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
			Lint(cfg),
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
