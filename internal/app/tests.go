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

func createTestFiles() string {
	// create a temp dir and file
	tempDir, err := os.MkdirTemp(WithProjectPath(".tests"), "test-data")
	if err != nil {
		log.Fatal(err)
	}

	// first: create a clusterfile
	yamlContent := `
---
version: 1
clusters:
  - name: unit-tests
    context: kind-kind
    envs:
      - name: web-apps
        location: addons.yaml
  - name: empty-cluster
    context: kind-kind-empty
    envs: []
`

	testClusterfile := path.Join(tempDir, "testdata.yaml")

	f, e := os.Create(testClusterfile)
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

	// then: create a test helmfile
	helmfile := `
---
version: 1
clusters:
  - name: unit-tests
    context: kind-kind
    envs:
      - name: web-apps
        location: addons.yaml
  - name: empty-cluster
    context: kind-kind-empty
    envs: []
`

	testfile := path.Join(tempDir, "addons.yaml")

	f, e = os.Create(testfile)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()
	if _, err := f.WriteString(helmfile); err != nil {
		log.Fatal(e)
	}
	if err := f.Close(); err != nil {
		log.Fatal(e)
	}

	return testClusterfile
}

func getTestCfg() types.Configuration {
	return types.Configuration{
		AppName:                "clusterfile-test",
		AppVersion:             "golang-test",
		AppUsage:               "Control the content of multiple k8s cluster via helmfile.",
		OverwrittenKubeContext: "kind-kind",
		SkipFlagParsing:        true,
		ClusterfileLocation:    createTestFiles(),
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
		SkipFlagParsing:        true,
		ClusterfileLocation:    createTestFiles(),
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
