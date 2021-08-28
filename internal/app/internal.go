package app

import (
	"encoding/json"
	"fmt"

	"github.com/la3mmchen/clusterfile/internal/types"
)

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
