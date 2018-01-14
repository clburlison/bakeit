package client

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Setup() {
	// Get the current node serial number
	// TODO: Limit this to window 7+. Will need to verify server lineup as well?
	clientConfig, err := GetConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create client config:\n%s\n", err)
		os.Exit(1)
	}

	// Write chef files
	err = ChefFiles(clientConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing chef files: %s\n", err)
		os.Exit(1)
	}

	// Install chef if required
	_, err = InstallChef()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bootstrap failed: %s\n", err)
		os.Exit(1)
	}

	// TODO: This is very opinionated. Should it be removed or
	// controlled with a config option?
	// Set the firstboot tag to ensure the firstboot runlist is used.
	err = ioutil.WriteFile("C:\\chef\\firstboot", []byte(""), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create firstboot file:%s\n", err)
	}

	// Run chef in a loop
	_, err = RunChef()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Chef failed to complete: %s\n", err)
		os.Exit(1)
	}
}
