package app

import (
	"github.com/la3mmchen/clusterfile/internal/types"
	"github.com/urfave/cli/v2"
)

func GetTestCfg() types.Configuration {
	return types.Configuration{
		AppName:             "clusterfile-test",
		AppVersion:          "golang-test",
		AppUsage:            "Control the content of multiple k8s cluster via helmfile.",
		ProjectPath:         GetProjectPath(),
		SkipFlagParsing:     true,
		ClusterfileLocation: "configs/clusterfile.yaml",
		AdditionalFlags: []cli.Flag{
			&cli.StringFlag{
				Name: "test.testlogfile",
			},
			&cli.StringFlag{
				Name: "test.paniconexit0",
			},
			&cli.StringFlag{
				Name: "test.v",
			},
		},
	}
}

func BootstrapTestApp() *cli.App {
	// construct an app for testing purposes
	cfg := GetTestCfg()

	return CreateApp(&cfg)
}

func BootstrapOfflineTestApp() *cli.App {
	// construct an app for testing purposes
	cfg := GetTestCfg()
	cfg.ActiveContext = ""

	return CreateApp(&cfg)
}
