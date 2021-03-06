package main

import (
	"embed"
	"fmt"
	"io/ioutil"
	"os"

	utils "github.com/pthomison/golang-utils"
	"github.com/spf13/cobra"
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

	installEnvs = []string{
		"K3S_NODE_NAME=alpha",
		"K3S_KUBECONFIG_OUTPUT=/home/pi/.kube/config",
		"K3S_KUBECONFIG_MODE=777",
	}
)

//go:embed k3s-install.sh
var installFile embed.FS

//go:embed manifests/*
var manifests embed.FS

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

	return executeAndAttach("bash", []string{"-c", tmpfileLocation}, installEnvs)
}

func coreload() {
	data, err := manifests.ReadFile("flux.yaml")
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

	return executeAndAttach("bash", []string{"-c", tmpfileLocation}, installEnvs)
}

func runInstall(cmd *cobra.Command, args []string) {
	utils.Check(install())
}

func runUninstall(cmd *cobra.Command, args []string) {
	utils.Check(executeAndAttach("k3s-uninstall.sh", nil, nil))
}

func runCoreLoad(cmd *cobra.Command, args []string) {
	utils.Check(executeAndAttach("k3s-uninstall.sh", nil, nil))
}

func runDummy(cmd *cobra.Command, args []string) {
	utils.Check(executeAndAttach("bash", []string{"-c", "for i in {1..10}; do echo ${i}; sleep 1; done"}, nil))
	fmt.Println("normal stdout")
	utils.Check(executeAndAttach("bash", []string{"-c", "for i in {20..30}; do echo ${i}; sleep 1; done"}, nil))
	fmt.Println("normal stdout 2")
}
