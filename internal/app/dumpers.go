package app

import (
	"encoding/json"
	"fmt"
)

func DumpMe(input interface{}) error { // TODO: is this the way to do a generic input?

	dumpMe, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		// how to catch this?!
		return nil
	}
	fmt.Printf("\n %s\n", string(dumpMe))

	return nil
}
