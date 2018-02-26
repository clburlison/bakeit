package chef

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/clburlison/bakeit/src/config"
	"github.com/clburlison/bakeit/src/download"
	"github.com/clburlison/bakeit/src/filetype"
	"github.com/groob/mackit/dmgutils"
	"github.com/groob/mackit/install/pkg"
)

// InstallChef will check to see if chef needs to be installed
// and if so install using the appropriate options per platform
func InstallChef() (bool, error) {
	// Check to see if chef is already installed and bail if on disk
	status := CheckChefInstall("/opt/chef/bin/chef-client")
	if status {
		return true, nil
	}

	// Download chef client
	// TODO: Instead of using the local directory path we should use
	// a temp file. The issue with this is we will then need to rename
	// dmg files and pkg files to have the proper extensions.
	// Apple's tools will not properly work on extensionless files.
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

	var (
		pkgPath       string
		dmgMountPoint string
	)

	// Check for file type
	ext, mime := filetype.Match(file)
	fmt.Printf("File type: %s. MIME: %s\n", ext, mime)

	// Handle dmgs, pkgs, and unsupported files
	if ext == "zlib" {
		mountpoints, err := dmgutils.MountDMG(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Mount failed: %v\n", err)
			return false, err
		}
		fmt.Printf("Mountpoints: %v \n", mountpoints)
		if len(mountpoints) > 1 {
			fmt.Fprintf(os.Stderr, "More than one mount path returned\n")
			return false, err
		}
		dmgMountPoint = mountpoints[0]
		packages := checkExt(".pkg", dmgMountPoint)
		fmt.Printf("Packages: %v \n", packages)
		if len(packages) > 1 {
			fmt.Fprintf(os.Stderr, "This dmg has more than one package\n")
			return false, err
		} else if packages == nil {
			fmt.Fprintf(os.Stderr, "Unable to locate packages in this dmg!\n")
			return false, err
		}
		pkgPath = filepath.Join(dmgMountPoint, packages[0])
		fmt.Printf("Package Path: %s\n", pkgPath)
	} else if ext == "xar" {
		pkgPath = file
	} else {
		fmt.Fprintf(os.Stderr, "Unsupported file type was downloaded!\n")
		return false, err
	}

	_, err = pkg.Install(pkgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Install failed: %v\n", err)
		return false, err
	}
	fmt.Printf("Chef installed sucessfully\n")

	// Unmount dmg files
	if ext == "zlib" {
		out, err := dmgutils.UnmountDMG(dmgMountPoint)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("Unmount result: %t\n", out)
	}

	// Remove Downloaded file
	os.Remove(file)
	return true, nil
}

// checkExt - Check a directory for files with a specific extension;
// return a listing of files found.
func checkExt(ext string, path string) []string {
	var files []string
	filepath.Walk(path, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	return files
}
