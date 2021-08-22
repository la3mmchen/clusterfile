package app

import (
	"log"
	"os"
	"path"

	"github.com/la3mmchen/clusterfile/internal/types"
	"github.com/urfave/cli/v2"
)

func BootstrapTestApp() *cli.App {
	// construct an app for testing purposes
	cfg := getTestCfg()
	return CreateApp(&cfg)
}

func BootstrapOfflineTestApp() *cli.App {
	// construct an app for testing purposes
	cfg := getBrokenTestCfg()

	return CreateApp(&cfg)
}

func createTestClusterfile() string {
	yamlContent := `
---
version: 1
clusters:
  - name: unit-tests
    context: kind-kind
    envs:
      - name: web-apps
        location: helmfile/addons.yaml
  - name: empty-cluster
    context: kind-kind-empty
    envs: []`

	// create a temp dir and file
	tempDir, err := os.MkdirTemp(path.Join(GetProjectPath(), ".tests"), "test-data")
	if err != nil {
		log.Fatal(err)
	}

	tempFile := path.Join(tempDir, "testdata.yaml")

	f, e := os.Create(tempFile)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()
	if _, err := f.WriteString(yamlContent); err != nil {
		log.Fatal(e)
	}
	if err := f.Close(); err != nil {
		log.Fatal(e)
	}

	return tempFile
}

func getTestCfg() types.Configuration {
	return types.Configuration{
		AppName:                "clusterfile-test",
		AppVersion:             "golang-test",
		AppUsage:               "Control the content of multiple k8s cluster via helmfile.",
		ProjectPath:            GetProjectPath(),
		OverwrittenKubeContext: "kind-kind",
		SkipFlagParsing:        true,
		ClusterfileLocation:    createTestClusterfile(),
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

func getBrokenTestCfg() types.Configuration {
	return types.Configuration{
		AppName:                "clusterfile-test",
		AppVersion:             "golang-test",
		AppUsage:               "Control the content of multiple k8s cluster via helmfile.",
		OverwrittenKubeContext: "broken-context",
		ProjectPath:            GetProjectPath(),
		SkipFlagParsing:        true,
		ClusterfileLocation:    createTestClusterfile(),
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
