package main

import (
	// "fmt"
	// "io/ioutil"
	"os"
	"os/exec"
	// "io/fs"
)

func executeAndAttach(c string, args []string, envs []string) error {
	cmd := exec.Command(c, args...)
	if envs != nil {
		cmd.Env = envs
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
