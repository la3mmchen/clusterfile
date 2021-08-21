package app

import (
	"fmt"
	"testing"
)

func TestPreloadCfg(t *testing.T) {
	cfg := GetTestCfg()

	// only test offline stuff
	cfg.PreflightConfig.Offline = true

	if err := PreloadCfg(&cfg); err != nil {
		t.Log(err)
		t.FailNow()
	}

}

func TestGetProjectPath(t *testing.T) {
	fmt.Printf("%v", GetProjectPath())

}
