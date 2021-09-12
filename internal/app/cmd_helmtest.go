package app

import (
	"fmt"

	"github.com/la3mmchen/clusterfile/internal/types"
	"github.com/urfave/cli/v2"
)

func HelmTest(cfg *types.Configuration) *cli.Command {
	cmd := cli.Command{
		Name:  "test",
		Usage: "run the test in all the helm charts.",
	}

	cmd.Action = func(c *cli.Context) error {

		err := PreloadCfg(cfg)
		if err != nil {
			return err
		}

		fmt.Printf("\nTesting environments for loaded kubernetes context [%v]. \n", cfg.ActiveContext)

		for i := range cfg.ActiveCluster.Envs {
			rc, err := RunWithRc(cfg.HelmfileExecutable, []string{"--file", cfg.ActiveCluster.Envs[i].Location, "test"}, false)
			if err != nil {
				return err
			}
			if rc == 0 {
				fmt.Printf("%v] test ok.", cfg.ActiveCluster.Envs[i].Name)
			} else {
				fmt.Printf("%v] testing failed.", cfg.ActiveCluster.Envs[i].Name)
			}

		}

		return nil
	}
	return &cmd
}
