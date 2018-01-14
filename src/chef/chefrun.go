package chef

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/clburlison/bakeit/src/config"
)

var logFile = config.FirstRunLogFile[runtime.GOOS]

// RunChef will call chef-client in a loop to setup a node
func RunChef() (status bool, err error) {
	// Set up the Chef Log directory
	logDir := config.FirstRunLogFile[runtime.GOOS]
	logDir = filepath.Dir(logDir)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, os.ModePerm)
	}

	retries := 3
	successes := 0
	for {
		fmt.Printf("Running Chef...\n")
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
			"logfile at %s", logFile)
		return false, err
	}

	fmt.Println("Bootstrap complete!")
	return true, nil
}

func callChefClient() (bool, error) {
	cmd := exec.Command(config.ChefClientExecPath[runtime.GOOS])
	// Create log file if it doesn't exist
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		_, err = os.Create(logFile)
	}
	// Open log file in append mode
	outfile, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0644)
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
