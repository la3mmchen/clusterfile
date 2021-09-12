package app

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/la3mmchen/clusterfile/internal/types"
	"gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// TODO: might be cool to add error handling
func WithProjectPath(s string) string {
	_, b, _, _ := runtime.Caller(0)
	//path := filepath.Join(filepath.Dir(b), "../..") // feels hacky but works
	path := filepath.Join(filepath.Dir(b), "../..", s) // feels hacky but works
	return path
}

func DiffEnv(cfg *types.Configuration, envfile string) (int, error) {
	// TODO: this returns 1 either if there is a diff in the envs or if the k8s cluster can not be reached
	return RunWithRc(cfg.HelmfileExecutable, []string{"--file", envfile, "diff", "--detailed-exitcode"}, true)
}

// PreloadCfg create parsed content into the global config struct
// it takes a pointer to the configuration and inserts certain
// values in there
// returns an error if preloading went wrong
func PreloadCfg(cfg *types.Configuration) error {
	var err error

	if len(cfg.OverwrittenKubeContext) == 0 {
		cfg.ActiveContext, err = GetActiveKubeContext()
	} else {
		cfg.ActiveContext = cfg.OverwrittenKubeContext
	}

	if err != nil {
		fmt.Printf("Error loading kube context: [%v] \n", err)
		return err
	}

	// parse clusterfile
	cfg.Clusterfile, err = ParseClusterfile(cfg)
	if err != nil {
		return err
	}

	if !SetActiveCluster(cfg) {
		return errors.New("can't find a definition for active kubernetes context in clusterfile")
	}

	err = ValidateEnvHelmfile(cfg)
	if err != nil {
		return err
	}

	return nil
}

// GetActiveKubeContext returns the current kubernetes context
// is loaded while running our app.
func GetActiveKubeContext() (string, error) {

	stdout, _, err := RunWithOutput("kubectl", []string{"config", "current-context"}) // i'm to stupid to do it with clientcmd

	if err != nil {
		return "", err
	}

	parsedContext := strings.Split(strings.TrimSuffix(stdout.String(), "\n"), "@")

	fmt.Printf("Parsed context from your env: [%v] \n", parsedContext[len(parsedContext)-1])

	return parsedContext[len(parsedContext)-1], nil
}

// ParseCLusterfile returns the parsed contents of the clusterfile
func ParseClusterfile(cfg *types.Configuration) (types.Clusterfile, error) {
	var tmpString string
	// check if cfg.Clusterfile is an absolute path
	fileinfo, err := os.Stat(cfg.ClusterfileLocation)

	if !os.IsNotExist(err) && !fileinfo.IsDir() {
		tmpString = cfg.ClusterfileLocation
	} else {
		tmpString = cfg.ClusterfileLocation
	}

	var clfl = types.Clusterfile{
		Location: tmpString,
	}

	f, err := os.Open(clfl.Location)

	if err != nil {
		return types.Clusterfile{}, err
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&clfl)
	if err != nil {
		return types.Clusterfile{}, err
	}

	return clfl, err
}

// SetActiveCluster determines the current cluster
// - copys the content to use  to a new position in config struct.
// - resolve possible path isseu in the defined sub-helmfle
func SetActiveCluster(cfg *types.Configuration) bool {
	var found = false

	// set active cluster
	for i := range cfg.Clusterfile.Clusters {
		if cfg.Clusterfile.Clusters[i].Context == cfg.ActiveContext {
			cfg.ActiveCluster = cfg.Clusterfile.Clusters[i]
			found = true
		}
	}

	// resolve relative or absolute paths in ActiveCluster.Envs.Locations
	for i := range cfg.ActiveCluster.Envs {
		if !filepath.IsAbs(cfg.ActiveCluster.Envs[i].Location) {
			cfg.ActiveCluster.Envs[i].Location = filepath.Join(filepath.Dir(cfg.ClusterfileLocation), cfg.ActiveCluster.Envs[i].Location)

		}
	}
	return found
}

// ValidateEnvHelmfile execute checks on ActiveCluster:
// - check the existence of the sub-helmfiles that are configured for the active cluster
// - drop envs that aren't selected
func ValidateEnvHelmfile(cfg *types.Configuration) error {
	for i := range cfg.ActiveCluster.Envs {

		// drop env if a specific env is selected via `--env`
		if len(cfg.EnvSelection) > 0 {
			if cfg.ActiveCluster.Envs[i].Name != cfg.EnvSelection {
				cfg.ActiveCluster.Envs = removeFromSliceByIndex(cfg.ActiveCluster.Envs, i)
				continue // next if the current env wasn't specified
			}
		}

		// check if the helmfile is present
		if _, err := os.Stat(cfg.ActiveCluster.Envs[i].Location); errors.Is(err, fs.ErrNotExist) {
			if cfg.Ignore {
				cfg.ActiveCluster.Envs = removeFromSliceByIndex(cfg.ActiveCluster.Envs, i)
			} else { // only return if we are not told to ignore via `--ignore`
				return fmt.Errorf("specific helmfile [%s] file not found", cfg.ActiveCluster.Envs[i].Location)
			}
		}
	}
	return nil
}

// CheckKubeConfig uses the current kubernetes context to
// test if the kubernetes cluster can be reached
func CheckKubeConfig(cfg *types.Configuration) error {
	var kubeconfig *string
	// use provided flag
	if len(cfg.OverwrittenKubeContext) > 0 {
		kubeconfig = &cfg.OverwrittenKubeContext
	} else {
		// check kubeconfig
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()
	}
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return err
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	_, err = clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	return nil
}
