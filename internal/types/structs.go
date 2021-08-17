package types

import "github.com/urfave/cli/v2"

// Configuration acts as central resource to save everything
type Configuration struct {
	AppName                string
	AppVersion             string
	AppUsage               string
	Debug                  string
	SkipFlagParsing        bool
	AdditionalFlags        []cli.Flag
	ClusterfileLocation    string
	Clusterfile            Clusterfile
	ActiveCluster          Cluster
	ActiveContext          string
	OverwrittenKubeContext string
	Helmfile               string
	HelmfileExecutable     string
	OutputDir              string
	PreflightConfig        Preflight
	TemplateConfig         Template
	BuildConfig            Build
	StatusConfig           Status
	Ignore                 bool
}

// Status represents options for the cli subcommand
type Status struct {
	Offline bool
}

// Build represents options for the cli subcommand
type Build struct {
	Stdout    bool
	GitCommit bool
}

// Template represents options for the cli subcommand
type Template struct {
	Stdout bool
}

// Preflight represents options for the cli subcommand
type Preflight struct {
	Offline bool
}

// Clusterfile contains the parsed clusterfile
type Clusterfile struct {
	Version  string    `yaml:"version"`
	Clusters []Cluster `yaml:"clusters"`
	Location string    // this is a meta information where the clusterfile is stored
}

type Cluster struct {
	Context string
	Envs    []Env `yaml:"envs"`
}

type Env struct {
	Name     string
	Location string
}

// --- delete afterwards maybe
type Release struct {
	Name      string
	Chart     Chart
	Installed string
	Version   string
	Namespace string
}

type Chart struct {
	Name    string
	Version string
}
