package commands

import (
	"fmt"

	"github.com/la3mmchen/clusterfile/internal/helpers"
	"github.com/la3mmchen/clusterfile/internal/types"
	"github.com/urfave/cli/v2"
)

func Diff(cfg *types.Configuration) *cli.Command {
	cmd := cli.Command{
		Name:  "diff",
		Usage: "diff the helmfile.",
	}

	cmd.Action = func(c *cli.Context) error {

		err := helpers.PreloadCfg(cfg)
		if err != nil {
			return err
		}

		for i := range cfg.ActiveCluster.Envs {
			rc, err := helpers.RunWithRc(cfg.HelmfileExecutable, []string{"--file", cfg.ActiveCluster.Envs[i].Location, "diff", "--detailed-exitcode"})
			if err != nil {
				return err
			}
			if rc == 0 {
				fmt.Printf("%v] state ok.", cfg.ActiveCluster.Envs[i].Name)
			} else {
				fmt.Printf("%v] state drifted. Sync needed.", cfg.ActiveCluster.Envs[i].Name)
			}

		}

		return nil
	}
	return &cmd
}
