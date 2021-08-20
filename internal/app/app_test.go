package app

import (
	"os"
	"strings"
	"testing"
)

func TestIfAppRuns(t *testing.T) {
	subCmd := []string{}
	app := BootstrapTestApp()

	args := os.Args[0:1]
	for i := range subCmd {
		args = append(args, subCmd[i])
	}

	if err := app.Run(args); err != nil {
		t.Fail()
		t.Logf("cli command [%v] failed. Error: %v", strings.Join(subCmd, ", "), err)
	}
}
