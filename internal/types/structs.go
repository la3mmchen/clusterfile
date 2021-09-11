package types

import "github.com/urfave/cli/v2"

// Configuration acts as central resource to save everything we have gotten or parsed
type Configuration struct {
	// input params - might be cool to move them to another struct
	AppName                string
	AppVersion             string
	AppUsage               string
	Debug                  string
	SkipFlagParsing        bool
	ClusterfileLocation    string
	OverwrittenKubeContext string
	Helmfile               string
	HelmfileExecutable     string
	OutputDir              string
	Ignore                 bool

	// link to other structs that contains input settings
	PreflightConfig Preflight
	TemplateConfig  Template
	BuildConfig     Build
	StatusConfig    Status

	// parsed content
	Clusterfile   Clusterfile
	ActiveCluster Cluster
	ActiveContext string

	// mixed stuff
	AdditionalFlags []cli.Flag
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

// Cluster contains the part of clusterfile that describes the cluster
type Cluster struct {
	Context string
	Envs    []Env `yaml:"envs"`
}

// Env contains the part of clusterfile that describes the env
type Env struct {
	Name     string
	Location string
}
