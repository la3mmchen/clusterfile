package app

import (
	"fmt"

	"github.com/la3mmchen/clusterfile/internal/types"
	"github.com/urfave/cli/v2"
)

func List(cfg *types.Configuration) *cli.Command {
	cmd := cli.Command{
		Name:  "list",
		Usage: "list all releases for the active cluster.",
	}

	cmd.Action = func(c *cli.Context) error {

		err := PreloadCfg(cfg)
		if err != nil {
			return err
		}

		for i := range cfg.ActiveCluster.Envs {
			stdout, _, err := RunWithOutput(cfg.HelmfileExecutable, []string{"--file", cfg.ActiveCluster.Envs[i].Location, "list"})
			if err != nil {
				return err
			}
			fmt.Printf("%s", stdout.String())
		}

		return nil
	}
	return &cmd
}
