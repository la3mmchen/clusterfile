package helpers

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/la3mmchen/clusterfile/internal/types"
	"gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func DiffEnv(cfg *types.Configuration, envfile string) (int, error) {
	return RunWithRc(cfg.HelmfileExecutable, []string{"--file", envfile, "diff", "--detailed-exitcode"}, true)
}

func PreloadCfg(cfg *types.Configuration) error {
	var err error

	cfg.ActiveContext, err = GetActiveKubeContext()

	if err != nil {
		fmt.Printf("Error loading kube context: [%v] \n", err)
		return err
	}

	// parse clusterfile
	cfg.Clusterfile, err = ParseClusterfile(cfg.ClusterfileLocation)
	if err != nil {
		return err
	}

	if !SetActiveCluster(cfg) {
		return errors.New("can't find a definition for active kubernetes context in clusterfile")
	}

	err = ValidateEnvHelmfile(cfg, cfg.Ignore)
	if err != nil {
		return err
	}
	return nil
}

func GetActiveKubeContext() (string, error) {

	stdout, _, err := RunWithOutput("kubectl", []string{"config", "current-context"}) // i'm to stupid to do it with clientcmd

	if err != nil {
		return "", err
	}

	parsedContext := strings.Split(strings.TrimSuffix(stdout.String(), "\n"), "@")

	fmt.Printf("Parsed context from your env: [%v] \n", parsedContext[len(parsedContext)-1])

	return parsedContext[len(parsedContext)-1], nil
}

func GetCommitSha() string {
	// TODO: implement me
	return "a6076f8" //TODO: just return a dev value
}

func CheckExecutable(cmd string) bool {

	_, err := exec.LookPath(cmd)
	return err == nil

}

func RunWithRc(prog string, args []string, silent bool) (int, error) {
	cmd := exec.Command(prog, args...)

	if !silent {
		fmt.Printf("Executing: [%v] \n", cmd)
	}

	err := cmd.Run()

	// probably the most stupid way to get the plan rc of the command ¯\_(ツ)_/¯
	exitCode := 0
	var e error
	if err != nil {
		exitCode, e = strconv.Atoi(strings.ReplaceAll(fmt.Sprintf("%v", err), "exit status ", ""))
	}

	if e != nil {
		return 42, e // return non zero integer
	}

	return exitCode, nil
}

func RunWithOutput(prog string, args []string) (bytes.Buffer, bytes.Buffer, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(prog, args...)

	fmt.Printf("Executing: [%v] \n", cmd)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return bytes.Buffer{}, stderr, err
	}

	return stdout, stderr, nil
}

func ParseClusterfile(clusterfile string) (types.Clusterfile, error) {
	var clfl = types.Clusterfile{
		Location: clusterfile,
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

func SetActiveCluster(cfg *types.Configuration) bool {
	var found = false

	for i := range cfg.Clusterfile.Clusters {
		if cfg.Clusterfile.Clusters[i].Context == cfg.ActiveContext {
			cfg.ActiveCluster = cfg.Clusterfile.Clusters[i]
			found = true
		}
	}
	return found
}

func ValidateEnvHelmfile(cfg *types.Configuration, ignore bool) error {

	for i := range cfg.ActiveCluster.Envs {
		if _, err := os.Stat(cfg.ActiveCluster.Envs[i].Location); errors.Is(err, fs.ErrNotExist) {
			if ignore {
				cfg.ActiveCluster.Envs = removeFromSliceByIndex(cfg.ActiveCluster.Envs, i)
			} else { // only return if we do not ignore the error
				return fmt.Errorf("specific helmfile [%s] file not found", cfg.ActiveCluster.Envs[i].Location)
			}
		}
	}

	return nil
}

func CheckKubeConfig() error {
	// check kubeconfig
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
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

func removeFromSliceByIndex(s []types.Env, index int) []types.Env {
	return append(s[:index], s[index+1:]...)
}
