package main

import (
	"fmt"
	"os"

	"github.com/la3mmchen/clusterfile/internal/app"
	"github.com/la3mmchen/clusterfile/internal/types"
)

var (
	// AppVersion Version of the app. Must be injected during the build.
	AppVersion string
	// Cfg types.Configuration
	Cfg types.Configuration
)

func main() {

	var cfg = types.Configuration{
		AppName:             "clusterfilectl",
		AppUsage:            "Control the content of multiple k8s cluster via helmfile.",
		AppVersion:          AppVersion,
		ClusterfileLocation: "data/clusterfile.yaml",
		ProjectPath:         app.GetProjectPath(),
		SkipFlagParsing:     false,
	}

	app := app.CreateApp(&cfg)

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error: %v \n", err)
		os.Exit(1)
	}
}
