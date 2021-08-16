package commands

import (
	"fmt"

	"github.com/la3mmchen/clusterfile/internal/helpers"
	"github.com/la3mmchen/clusterfile/internal/types"
	"github.com/urfave/cli/v2"
)

func Lint(cfg *types.Configuration) *cli.Command {
	cmd := cli.Command{
		Name:  "lint",
		Usage: "lint the helmfile.",
	}

	cmd.Action = func(c *cli.Context) error {

		err := helpers.PreloadCfg(cfg)
		if err != nil {
			return err
		}

		for i := range cfg.ActiveCluster.Envs {
			stdout, _, err := helpers.RunWithOutput(cfg.HelmfileExecutable, []string{"--file", cfg.ActiveCluster.Envs[i].Location, "lint"})
			if err != nil {
				return err
			}
			fmt.Printf("%s", stdout.String())
		}

		return nil
	}
	return &cmd
}
