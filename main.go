package main

import (
	"embed"
	"fmt"
	utils "github.com/pthomison/golang-utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	// "os/exec"
	// "io/fs"
)

var (
	// message string
	// name    string

	rootCmd = &cobra.Command{
		Use:   "k3s-helper",
		Short: "k3s-helper",
	}

	installCmd = &cobra.Command{
		Use:   "install",
		Short: "install",
		Run:   runInstall,
	}

	uninstallCmd = &cobra.Command{
		Use:   "uninstall",
		Short: "uninstall",
		Run:   runUninstall,
	}

	dummyCmd = &cobra.Command{
		Use:   "dummy",
		Short: "dummy",
		Run:   runDummy,
	}
)

//go:embed k3s-install.sh
var installFile embed.FS

func init() {
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(uninstallCmd)
	rootCmd.AddCommand(dummyCmd)
}

func main() {

	// rootCmd.PersistentFlags().StringVarP(&message, "message", "m", "hello world", "message the program will output")
	// rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "patrick", "name the program will output to")

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

	return executeAndAttach("bash", "-c", tmpfileLocation)
}

func runInstall(cmd *cobra.Command, args []string) {
	utils.Check(install())
}

func runUninstall(cmd *cobra.Command, args []string) {
	utils.Check(executeAndAttach("k3s-uninstall.sh"))
}

func runDummy(cmd *cobra.Command, args []string) {
	utils.Check(executeAndAttach("bash", "-c", "for i in {1..10}; do echo ${i}; sleep 1; done"))
	fmt.Println("normal stdout")
	utils.Check(executeAndAttach("bash", "-c", "for i in {20..30}; do echo ${i}; sleep 1; done"))
	fmt.Println("normal stdout 2")
}
