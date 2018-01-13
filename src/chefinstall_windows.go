package client

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// https://docs.chef.io/install_bootstrap.html#powershell-user-data

// InstallChef will check to see if chef needs to be installed
// and if so install using the appropriate options per platform
func InstallChef() (bool, error) {
	// Check to see if chef is already installed and bail if on disk
	status := CheckChefInstall("C:\\opscode\\chef\\bin\\chef-client")
	if status {
		return true, nil
	}

	// Download chef client
	// TODO: Instead of using the local directory path we should use
	// a temp file. Verify windows can install a msi without the
	// extension
	file, err := Download(GetChefURL(), ".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading file: %s\n", err)
	}
	fmt.Printf("Downloaded file: %s\n", file)

	fmt.Printf("Installing chef...\n")
	_, err = installMSI("C:\\Windows\\Temp\\chef-log.txt",
		file,
		`ADDLOCAL="ChefClientFeature,ChefSchTaskFeature,ChefPSModuleFeature"`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Install failed: %v\n", err)
		return false, err
	}
	fmt.Printf("Chef installed sucessfully\n")

	// Remove downloaded file
	os.Remove(file)
	return true, nil
}

// installMSI- Install an msi file
// TODO: extraArgs should take an unlimited array instead of a single
// string. We should then unwrap this before passing to exec
func installMSI(logFile string, msiFile string, extraArgs string) (string, error) {
	cmdName := "msiexec.exe"
	// TODO: extraArgs is not being passed sucessfully to cmdArgs, as such
	// it is currently excluded.
	cmdArgs := []string{"/qn", "/lv", logFile, "/i", msiFile}
	cmd := exec.Command(cmdName, cmdArgs...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
