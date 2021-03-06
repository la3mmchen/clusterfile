package types

import "github.com/urfave/cli/v2"

// Configuration acts as central resource to save everything we have gotten or parsed
type Configuration struct {
	// input params - might be cool to move them to another struct
	AppName                string
	AppVersion             string
	AppUsage               string
	ClusterfileLocation    string
	Debug                  string
	EnvSelection           string
	Ignore                 bool
	Helmfile               string
	HelmfileExecutable     string
	OverwrittenKubeContext string
	OutputDir              string
	SkipFlagParsing        bool
	Offline                bool

	// link to other structs that contains input settings
	TemplateConfig Template
	BuildConfig    Build

	// parsed content
	Clusterfile   Clusterfile
	ActiveCluster Cluster
	ActiveContext string

	// mixed stuff
	AdditionalFlags []cli.Flag
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

// Clusterfile contains the parsed clusterfile
type Clusterfile struct {
	Version  string    `yaml:"version"`
	Clusters []Cluster `yaml:"clusters"`
	Location string    // this is a meta information where the clusterfile is stored
}

// Cluster contains the part of clusterfile that describes the cluster
type Cluster struct {
	Context  string
	Releases []Release `yaml:"releases,omitempty"`
	Envs     []Env     `yaml:"envs"`
}

// Env contains the part of clusterfile that describes the env
type Env struct {
	Name     string
	Location string
}

// Release contains the part of clusterfile that describes releases
type Release struct {
	Name    string
	Version string
}
