package main

import (
	// "fmt"
	// "io/ioutil"
	"os"
	"os/exec"
	// "io/fs"
)

func executeAndAttach(c string, args ...string) error {
	cmd := exec.Command(c, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
