package setup

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"

	"github.com/clburlison/bakeit/src/chef"
	"github.com/clburlison/bakeit/src/config"
	"github.com/clburlison/bakeit/src/node"
)

// Setup is the main platform specific function that is called
// to setup a chef node.
func Setup() {
	// Only run on supported OS versions
	info, err := node.GetNodeInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error obtaining OS Version: %s\n", err)
		os.Exit(1)
	}
	osVer := info.ProductVersion
	majorVer := strings.Split(osVer, ".")[1]
	i, err := strconv.Atoi(majorVer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error obtaining OS Version: %s\n", err)
		os.Exit(1)
	}
	if i < 10 {
		fmt.Fprintf(os.Stderr, "'%s' is no longer supported. This machine "+
			"must be upgraded to install Chef.\n", osVer)
		os.Exit(1)
	}

	// Run with elevated permissions
	user, _ := user.Current()
	if user.Uid != "0" {
		fmt.Println("Please run as root!")
		os.Exit(1)
	}

	clientConfig, err := chef.GetConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create client config:\n%s\n", err)
		os.Exit(1)
	}

	// If current user is 'config.UserShortName', abort
	configUser := config.UserShortName
	if configUser == consoleUser() {
		fmt.Fprintf(os.Stderr, "Create a new account not named "+
			"'%s'!\nChef will create a '%s' account for you.", configUser, configUser)
		os.Exit(1)
	}

	// Write chef files
	err = node.ChefFiles(clientConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing chef files: %s ", err)
		os.Exit(1)
	}

	// Install chef if required
	_, err = chef.InstallChef()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bootstrap failed: %s\n", err)
		os.Exit(1)
	}

	// TODO: This is very opinionated. Should it be removed or
	// controlled with a config option?
	// Set the firstboot tag to ensure the firstboot runlist is used.
	err = ioutil.WriteFile("/etc/chef/firstboot", []byte(""), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create firstboot file:%s\n", err)
	}

	// Run chef in a loop
	_, err = chef.RunChef()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Chef failed to complete: %s\n", err)
		os.Exit(1)
	}
}

func consoleUser() string {
	cmd := exec.Command("/usr/bin/stat", "-f%Su", "/dev/console")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("consoleUser:", err)
	}
	return strings.TrimSpace(out.String())
}
