package client

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"regexp"

	"github.com/groob/mackit/dmgutils"
	"github.com/groob/mackit/install/pkg"
)

func Setup() {
	user, _ := user.Current()
	if user.Uid != "0" {
		fmt.Println("Please run as root!")
		os.Exit(1)
	}
	file, err := Download(".", "https://packages.chef.io/files/stable/chef/13.6.4/mac_os_x/10.13/chef-13.6.4-1.dmg", true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading file: %s\n", err)
		os.Exit(1)
	}

	pkgPath := ""
	dmgMountPoint := ""

	// Check for file type
	ext, mime := Match(file)
	fmt.Printf("File type: %s. MIME: %s\n", ext, mime)

	// Handle dmgs, pkgs, and unsupported files
	if ext == "zlib" {
		mountpoints, err := dmgutils.MountDMG(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Mount failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Mountpoints: %v \n", mountpoints)
		if len(mountpoints) > 1 {
			fmt.Fprintf(os.Stderr, "More than one mount path returned\n")
			os.Exit(1)
		}
		dmgMountPoint = mountpoints[0]
		packages := checkExt(".pkg", dmgMountPoint)
		fmt.Printf("Packages: %v \n", packages)
		if len(packages) > 1 {
			fmt.Fprintf(os.Stderr, "This dmg has more than one package\n")
			os.Exit(1)
		} else if packages == nil {
			fmt.Fprintf(os.Stderr, "Unable to locate packages in this dmg!\n")
			os.Exit(1)
		}
		pkgPath = filepath.Join(dmgMountPoint, packages[0])
		fmt.Printf("Package Path: %s\n", pkgPath)
	} else if ext == "xar" {
		pkgPath = file
	} else {
		fmt.Fprintf(os.Stderr, "Unsupported file type was downloaded!\n")
		os.Exit(1)
	}

	restart, err := pkg.Install(pkgPath)
	fmt.Printf("Restart value is: %v \n", restart)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Install failed: %v\n", err)
	}

	// Unmount dmg files
	if ext == "zlib" {
		out, err := dmgutils.UnmountDMG(dmgMountPoint)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("Unmount result: %t\n", out)
	}
}

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
