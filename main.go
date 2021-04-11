package main

import (
	"embed"
	"fmt"
	utils "github.com/pthomison/golang-utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	// "io/fs"
)

var (
	message string
	name    string

	rootCmd = &cobra.Command{
		Use:   "go-helloworld",
		Short: "go-helloworld",
		Run:   run,
	}
)

//go:embed k3s-install.sh
var installFile embed.FS

func main() {

	rootCmd.PersistentFlags().StringVarP(&message, "message", "m", "hello world", "message the program will output")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "patrick", "name the program will output to")

	err := rootCmd.Execute()

	utils.Check(err)
}

func install() error {
	data, err := installFile.ReadFile("k3s-install.sh")
	if err != nil {
		return err
	}

	tmpfile, err := ioutil.TempFile("", "k3s-install-script")
	if err != nil {
		return err
	}

	tmpfileLocation := tmpfile.Name()

	defer os.Remove(tmpfileLocation)

	err = tmpfile.Chmod(0700)
	if err != nil {
		return err
	}

	if _, err := tmpfile.Write(data); err != nil {
		return err
	}
	if err := tmpfile.Close(); err != nil {
		return err
	}

	cmd := exec.Command("bash", "-c", tmpfileLocation)

	out, err := cmd.Output()
	if err != nil {
		return err
	}

	fmt.Printf("output: %+v\n", string(out))

	return err

	// return nil
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println(message + ", " + name)
	utils.Check(install())

}