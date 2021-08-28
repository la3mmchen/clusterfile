package app

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

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
