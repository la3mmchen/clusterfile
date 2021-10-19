package types

/*
	Describes a legacy structure
*/

type ClusterfileLegacy struct {
	Clusters map[string]ClusterLegacy `yaml:"clusters"`
	Location string
}

type ClusterLegacy struct {
	Envs       []string          `yaml:"envs,omitempty"`
	Overwrites map[string]EnvMap `yaml:"overwrites,omitempty"`
}

type EnvMap struct {
}
