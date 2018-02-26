package chef

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/clburlison/bakeit/src/config"
	"github.com/clburlison/bakeit/src/download"
)

// https://docs.chef.io/install_bootstrap.html#powershell-user-data
// https://github.com/gorillalabs/go-powershell

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
	file, err := download.Download(download.GetChefURL(), ".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading file: %s\n", err)
		return false, err
	}

	// Verify the downloaded file matches the expected checksum value
	checksum := config.ChefClientURLChecksum[runtime.GOOS]
	if checksum != "" {
		hash, err := download.CheckHash(file, config.ChefClientURLChecksum[runtime.GOOS])
		if err != nil {
			return false, fmt.Errorf("InstallChef: Unable to verify hash value: %s", err)
		}
		if hash != true {
			return false, fmt.Errorf("InstallChef: Hash value does not match")
		}
	}

	fmt.Printf("Installing chef...\n")
	_, err = installMSI("C:\\Windows\\Temp\\chef-log.txt",
		file,
		`ADDLOCAL=ChefClientFeature,ChefSchTaskFeature,ChefPSModuleFeature`)
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
	cmdArgs := []string{"/qn", "/lv", logFile, "/i", msiFile, extraArgs}
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
