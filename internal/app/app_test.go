package app

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/la3mmchen/clusterfile/internal/types"
)

func TestCreateApp(t *testing.T) {
	var cfg types.Configuration
	app := CreateApp(&cfg)

	if _, err := app.ToMan(); err != nil {
		fmt.Printf("Error: %v \n", err)
		os.Exit(1)
	}
}

func TestIfAppRuns(t *testing.T) {
	subCmd := []string{}
	app := BootstrapTestApp()

	args := os.Args[0:1]
	for i := range subCmd {
		args = append(args, subCmd[i])
	}

	if err := app.Run(args); err != nil {
		t.Logf("cli command [%v] failed. Error: %v", strings.Join(subCmd, ", "), err)
		t.FailNow()
	}
}
