package app

import (
	"fmt"
	"os"

	"github.com/la3mmchen/clusterfile/internal/types"
	"github.com/urfave/cli/v2"
)

func Template(cfg *types.Configuration) *cli.Command {
	cmd := cli.Command{
		Name:  "template",
		Usage: "template releases from state file(s)",
	}

	cmd.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "stdout",
			Value:       false,
			Destination: &cfg.TemplateConfig.Stdout,
			DefaultText: "default: false, do print templated value to stdout as well",
			Usage:       "Toggle output to stdout",
		},
	}

	cmd.Action = func(c *cli.Context) error {

		err := PreloadCfg(cfg)
		if err != nil {
			return err
		}

		for i := range cfg.ActiveCluster.Envs {
			stdout, stderr, err := RunWithOutput(cfg.HelmfileExecutable, []string{"--file", cfg.ActiveCluster.Envs[i].Location, "template"})
			if err != nil {
				return err
			}

			// as helmfile template prints status to stderr we need print stderr too ¯\_(ツ)_/¯
			fmt.Printf("%s", stderr.String())

			// create a randome file to write the templated value
			f, e := os.CreateTemp(cfg.OutputDir, "rendered-manifests-*.yaml")
			if e != nil {
				return e
			}
			defer f.Close()

			if _, err := f.Write(stdout.Bytes()); err != nil {
				return err
			}
			if err := f.Close(); err != nil {
				return err
			}
			fmt.Printf("Wrote generated manifests to: [%v] \n", f.Name())

			if cfg.TemplateConfig.Stdout {
				fmt.Printf("%s", stdout.String())
			}

		}

		return nil
	}
	return &cmd
}
