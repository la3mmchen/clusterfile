package app

import (
	"fmt"
	"testing"
)

func TestWithProjectPath(t *testing.T) {
	fmt.Printf("%v", WithProjectPath("file"))

}

func TestPreloadCfg(t *testing.T) {
	cfg := getTestCfg()

	// only test offline stuff
	cfg.Offline = true

	if err := PreloadCfg(&cfg); err != nil {
		t.Log(err)
		t.FailNow()
	}

}
