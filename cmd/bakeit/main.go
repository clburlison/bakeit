package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"time"

	"github.com/cavaliercoder/grab"
	bakeitClient "github.com/clburlison/bakeit/client"
	"github.com/groob/mackit/dmgutils"
	"github.com/groob/mackit/install/pkg"
)

func main() {
	config := bakeitClient.Config()
	fmt.Printf("%s", config)
	user, err := user.Current()
	if user.Uid != "0" {
		fmt.Println("Please run as root!")
		os.Exit(1)
	}

	// create client
	c := grab.NewClient()
	req, _ := grab.NewRequest(".", "https://packages.chef.io/files/stable/chef/13.6.4/mac_os_x/10.13/chef-13.6.4-1.dmg")
	// req, _ := grab.NewRequest("./munki.pkg", "https://github.com/munki/munki/releases/download/v3.1.1/munkitools-3.1.1.3447.pkg")

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := c.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)

	pkgPath := ""
	dmgMountPoint := ""

	// Check for file type
	ext, mime := bakeitClient.Match(resp.Filename)
	fmt.Printf("File type: %s. MIME: %s\n", ext, mime)

	// Handle dmgs, pkgs, and unsupported files
	if ext == "zlib" {
		mountpoints, err := dmgutils.MountDMG(resp.Filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Mount failed: %v\n", err)
			os.Exit(1)
		}
		dmgMountPoint = mountpoints[0] // TODO: Can I assume 0 is correct?
		packages := checkExt(".pkg", dmgMountPoint)
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
		pkgPath = resp.Filename
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
