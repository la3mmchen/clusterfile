package app

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"

	"github.com/la3mmchen/clusterfile/internal/types"
	"github.com/urfave/cli/v2"
)

func Build(cfg *types.Configuration) *cli.Command {
	cmd := cli.Command{
		Name:            "build",
		Usage:           "generate a helmfile state",
		SkipFlagParsing: cfg.SkipFlagParsing,
	}

	cmd.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "stdout",
			Value:       false,
			Destination: &cfg.BuildConfig.Stdout,
			DefaultText: "default: false, do print buildd value to stdout as well.",
			Usage:       "Toggle output to stdout",
		},
		&cli.BoolFlag{
			Name:        "git-commit",
			Value:       false,
			Destination: &cfg.BuildConfig.GitCommit,
			DefaultText: "default: false, use git short commit id for rendered file.",
			Usage:       "Use git commit sha for naming rendered files",
		},
	}

	cmd.Action = func(c *cli.Context) error {

		err := PreloadCfg(cfg)
		if err != nil {
			return err
		}

		for i := range cfg.ActiveCluster.Envs {

			stdout, stderr, err := RunWithOutput(cfg.HelmfileExecutable, []string{"--file", cfg.ActiveCluster.Envs[i].Location, "build"})
			if err != nil {
				return err
			}
			// as helmfile build prints status to stderr we need print stderr too ¯\_(ツ)_/¯
			fmt.Printf("%s", stderr.String())

			// create a randome file to write the build value
			var name string
			if cfg.BuildConfig.GitCommit {
				name = fmt.Sprintf("%s.yaml", getCommitSha())
			} else {
				rand.Seed(time.Now().UnixNano())
				chars := []rune("abcdefghijklmnopqrstuvwxyz" +
					"0123456789")
				length := 8
				var b strings.Builder
				for i := 0; i < length; i++ {
					b.WriteRune(chars[rand.Intn(len(chars))])
				}
				name = fmt.Sprintf("rendered-helmfile-%s.yaml", b.String())
			}

			// make sure the output dir is present

			if _, err := os.Stat(cfg.OutputDir); !os.IsNotExist(err) {
				_ = os.MkdirAll(cfg.OutputDir, 0740)
			}

			f, e := os.Create(path.Join(cfg.OutputDir, name))
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

			if cfg.BuildConfig.Stdout {
				fmt.Printf("%s", stdout.String())
			}
		}
		return nil
	}
	return &cmd
}
