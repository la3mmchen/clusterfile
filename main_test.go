package main

import (
	"os"
	"strings"
	"testing"

	"github.com/la3mmchen/clusterfile/internal/app"
)

// TestAppWithoutKubeContext
// - TODO: test how app behaves without working kubeconfig
// - TODO: how do the app behaves without kubeconfig file
func TestAppWithoutKubeContext(t *testing.T) {
	subCmd := []string{"preflight"}
	app := app.BootstrapOfflineTestApp()

	args := os.Args[0:1]
	for i := range subCmd {
		args = append(args, subCmd[i])
	}

	// fail if the app does not fail
	if err := app.Run(args); err == nil {
		t.Fail()
		t.Logf("cli command [%v] failed. Expected error, got none.\n ", strings.Join(subCmd, ", "))
	}
}

func TestAppRun(t *testing.T) {
	subCmd := []string{}
	app := app.BootstrapTestApp()

	args := os.Args[0:1]
	for i := range subCmd {
		args = append(args, subCmd[i])
	}

	if err := app.Run(args); err != nil {
		t.Fail()
		t.Logf("cli command [%v] failed. Error: %v", strings.Join(subCmd, ", "), err)
	}
}

func TestSubcmds(t *testing.T) {
	cases := map[string][]string{
		"build":             {"build"},
		"dump":              {"dump"},
		"list":              {"list"},
		"lint":              {"lint"},
		"preflight":         {"preflight"},
		"preflight-offline": {"preflight", "--offline"},
		"status":            {"status"},
		"status-offline":    {"status", "--offline"},
	}
	args := os.Args[0:1]
	for testcase, subcmds := range cases {
		// create a new test app
		app := app.BootstrapTestApp()

		for i := range subcmds {
			args = append(args, subcmds[i])
		}

		if err := app.Run(args); err != nil {
			t.Logf("SubCmd [%v]: cli command [%v] failed. Error: %v", testcase, strings.Join(subcmds, ", "), err)
			t.FailNow()
		}
	}
}
