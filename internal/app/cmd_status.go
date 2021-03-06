package app

import (
	"fmt"

	"github.com/la3mmchen/clusterfile/internal/types"
	"github.com/urfave/cli/v2"
)

func Status(cfg *types.Configuration) *cli.Command {
	cmd := cli.Command{
		Name:  "status",
		Usage: "check the status of the defined envs for the loaded context.",
	}

	cmd.Action = func(c *cli.Context) error {

		err := PreloadCfg(cfg)
		if err != nil {
			return err
		}

		if len(cfg.ActiveCluster.Envs) > 0 {
			fmt.Printf("\nConfigured environments for loaded kubernetes context [%v]. \n", cfg.ActiveContext)
			envs := make(map[string]string)

			for i := range cfg.ActiveCluster.Envs {
				envs[cfg.ActiveCluster.Envs[i].Name] = "state not identified."
				if !cfg.Offline {
					rc, err := DiffEnv(cfg, cfg.ActiveCluster.Envs[i].Location)
					if err != nil {
						return err
					}
					if rc == 0 {
						envs[cfg.ActiveCluster.Envs[i].Name] = "state ok"
					} else {
						envs[cfg.ActiveCluster.Envs[i].Name] = "state drifted. Sync needed"
					}
				}
			}
			for k, v := range envs {
				fmt.Printf("- %v [%v] \n", k, v)
			}
		} else {
			fmt.Printf("\nNo environments defined for this kubernetescontext [%v]. \n", cfg.ActiveContext)
		}

		return nil
	}

	return &cmd
}
