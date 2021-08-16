package commands

import (
	"github.com/la3mmchen/clusterfile/internal/helpers"
	"github.com/la3mmchen/clusterfile/internal/types"
	"github.com/urfave/cli/v2"
)

func Dump(cfg *types.Configuration) *cli.Command {
	cmd := cli.Command{
		Name:  "dump",
		Usage: "dump the parsed information.",
	}

	cmd.Action = func(c *cli.Context) error {

		err := helpers.PreloadCfg(cfg)
		if err != nil {
			return err
		}

		types.DumpMe(cfg)

		return nil
	}
	return &cmd
}
