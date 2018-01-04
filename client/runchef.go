package client

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	logDir          = "/Library/Chef/Logs"
	firstRunLogFile = "/Library/Chef/Logs/first_chef_run.log"
)

// RunChef will call chef-client in a loop to setup a node
func RunChef() (status bool, err error) {
	// Set up the Chef Log directory
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, os.ModePerm)
	}

	retries := 3
	successes := 0
	for {
		fmt.Printf("Running Chef %d/%d\n", successes+1, retries)
		success, _ := callChefClient()
		if success {
			fmt.Printf("Chef ran successfully...\n")
			successes++
		} else {
			fmt.Println("Chef run failed, retrying...")
			retries--
		}
		if successes == 2 {
			break
		}
		if retries == 0 {
			break
		}
	}
	if successes < 2 {
		err := fmt.Errorf("Chef failed to run, please send "+
			"logfile at %s", firstRunLogFile)
		return false, err
	}

	fmt.Println("Bootstrap complete!")
	return true, nil
}

func callChefClient() (bool, error) {
	cmd := exec.Command("/usr/local/bin/chef-client")
	// Create log file if it doesn't exist
	if _, err := os.Stat(firstRunLogFile); os.IsNotExist(err) {
		_, err = os.Create(firstRunLogFile)
	}
	// Open log file in append mode
	outfile, err := os.OpenFile(firstRunLogFile, os.O_APPEND|os.O_WRONLY, 0644)
	defer outfile.Close()
	if err != nil {
		return false, err
	}
	cmd.Stdout = outfile
	err = cmd.Run()
	if err != nil {
		return false, err
	}
	return true, nil
}
