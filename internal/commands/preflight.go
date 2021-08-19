package commands

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/la3mmchen/clusterfile/internal/helpers"
	"github.com/la3mmchen/clusterfile/internal/types"
	"github.com/urfave/cli/v2"
)

func Preflight(cfg *types.Configuration) *cli.Command {
	cmd := cli.Command{
		Name:  "preflight",
		Usage: "Execute prefligth checks to assert that the app works in your environment.",
	}

	cmd.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "offline",
			Value:       false,
			Destination: &cfg.PreflightConfig.Offline,
			DefaultText: "default: false, do connect to configured kubernetes cluster",
			Usage:       "skip kubernetes connect",
		},
	}

	cmd.Action = func(c *cli.Context) error {

		var prefixText = "Preflight check: "

		// check kubeconfig
		if !cfg.PreflightConfig.Offline {
			fmt.Printf("%schecking %s \n", prefixText, cfg.ClusterfileLocation)
			err := helpers.CheckKubeConfig()
			if err != nil {
				return err
			}
			fmt.Printf(" ok. \n")
		}

		// parse clusterfile
		fmt.Printf("%schecking %s.", prefixText, cfg.ClusterfileLocation)
		_, err := helpers.ParseClusterfile(filepath.Join(cfg.ProjectPath, cfg.ClusterfileLocation))
		if err != nil {
			fmt.Printf(err.Error())
			return err
		}
		fmt.Printf(" ok. \n")

		// check if helmfile executable is present
		fmt.Printf("%shelmfile executable.", prefixText)
		if !helpers.CheckExecutable(cfg.HelmfileExecutable) {
			return errors.New("executable not found in PATH")
		}
		fmt.Printf(" ok. \n")

		return nil
	}
	return &cmd
}
