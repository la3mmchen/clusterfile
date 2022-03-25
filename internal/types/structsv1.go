package types

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
