package client

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// CheckChefInstall will return true/false if chef is installed
func CheckChefInstall(clientPath string) bool {
	if _, err := os.Stat(clientPath); os.IsNotExist(err) {
		return false
	}
	// We aren't using the version at this time
	status, _ := chefVersion(clientPath)
	if status {
		return true
	}
	return false
}

func chefVersion(clientPath string) (bool, string) {
	cmd := exec.Command(clientPath, "-v")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error with chef-client: %s\n", err)
		return false, ""
	}
	if strings.Contains(out.String(), ": ") {
		ver := strings.Split(out.String(), ": ")[1]
		ver = strings.TrimSpace(ver)
		fmt.Printf("Chef %s installed\n", ver)
		return true, ver
	}
	fmt.Printf("Chef installed, but we could not get the version\n")
	return true, "UNKNOWN"
}
