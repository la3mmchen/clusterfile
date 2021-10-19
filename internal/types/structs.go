package types

import "github.com/urfave/cli/v2"

// Configuration acts as central resource to save everything we have gotten or parsed
type Configuration struct {
	// input params - might be cool to move them to another struct
	AppName                string
	AppVersion             string
	AppUsage               string
	ClusterfileLocation    string
	ClusterfileVersion     string
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
	Clusterfile       Clusterfile
	ClusterfileLegacy ClusterfileLegacy
	ActiveCluster     Cluster
	ActiveContext     string

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
