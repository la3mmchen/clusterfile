package app

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/la3mmchen/clusterfile/internal/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// checkKubeConfig uses the current kubernetes context to
// test if the kubernetes cluster can be reached
func checkKubeConfig(cfg *types.Configuration) error {

	tmpFlags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var kubeconfig *string
	// use provided flag
	if len(cfg.OverwrittenKubeContext) > 0 {
		kubeconfig = &cfg.OverwrittenKubeContext
	} else {
		// check kubeconfig
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = tmpFlags.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = tmpFlags.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
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

func removeFromSliceByIndex(s []types.Env, index int) []types.Env {
	return append(s[:index], s[index+1:]...)
}

func dumpMe(input interface{}) error { // TODO: is this the way to do a generic input?

	dumpMe, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		// how to catch this?!
		return nil
	}
	fmt.Printf("\n %s\n", string(dumpMe))

	return nil
}

func getCommitSha() string {
	// TODO: implement me
	return "a6076f8" //TODO: just return a dev value
}

// GetActiveKubeContext returns the current kubernetes context
// is loaded while running our app.
func getActiveKubeContext() (string, error) {

	stdout, _, err := RunWithOutput("kubectl", []string{"config", "current-context"}) // i'm to stupid to do it with clientcmd

	if err != nil {
		return "", err
	}

	parsedContext := strings.Split(strings.TrimSuffix(stdout.String(), "\n"), "@")

	fmt.Printf("Parsed context from your env: [%v] \n", parsedContext[len(parsedContext)-1])

	return parsedContext[len(parsedContext)-1], nil
}
