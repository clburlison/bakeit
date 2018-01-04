package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/clburlison/bakeit/client/config"
	"github.com/groob/mackit/dmgutils"
	"github.com/groob/mackit/install/pkg"
)

// Setup is the main platform specific function that is called
// to setup a chef node.
func Setup() {
	// Only run on supported OS versions
	info, err := GetMacInfo()
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

	// TODO: Start: Move this to the client.go file as a separate function
	// Get the current node serial number
	serial := GetSerialNumber()
	fmt.Printf("Current serial number is: %s\n", serial)

	// Build client.rb from config and template
	settings := Settings{
		config.ChefClientLogLevel,
		config.ChefClientLogLocation,
		config.ChefClientValidationClientName,
		config.ChefClientValidationKey,
		config.ChefClientChefServerURL,
		config.ChefClientJSONAttribs,
		config.ChefClientSSLVerifyMode,
		config.ChefClientLocalKeyGeneration,
		config.ChefClientRestTimeout,
		config.ChefClientHTTPRetryCount,
		config.ChefClientNoLazyLoad,
		config.ChefClientOhaiDirectory["darwin"],
		config.ChefClientOhaiDisabledPlugins["darwin"],
		serial}
	clientConfig, err := Config(settings)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create client config:\n%s\n", err)
		os.Exit(1)
	}
	// TODO: End

	// Run with elevated permissions
	user, _ := user.Current()
	if user.Uid != "0" {
		fmt.Println("Please run as root!")
		os.Exit(1)
	}

	// If current user is 'config.UserShortName', abort
	configUser := config.UserShortName
	if configUser == consoleUser() {
		fmt.Fprintf(os.Stderr, "Create a new account not named "+
			"'%s'!\nChef will create a '%s' account for you.", configUser, configUser)
		os.Exit(1)
	}

	// Remove old chef files before running
	if config.Force {
		clientCert := "/etc/chef/client.pem"
		if _, err := os.Stat(clientCert); os.IsExist(err) {
			os.Remove(clientCert)
		}
		ohaiPlugins := "/etc/chef/ohai_plugins/"
		if _, err := os.Stat(ohaiPlugins); os.IsExist(err) {
			os.Remove(ohaiPlugins)
		}
	}

	// Write chef config files
	if _, err := os.Stat("/etc/chef"); os.IsNotExist(err) {
		os.MkdirAll("/etc/chef", os.ModePerm)
	}
	err = ioutil.WriteFile("/etc/chef/client.rb", []byte(clientConfig), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to write client config:%s\n", err)
	}
	json := config.ChefClientRunListJSON["darwin"]
	err = ioutil.WriteFile("/etc/chef/run-list.json", []byte(json), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error formating JSON run list:%s\n", err)
	}
	orgCrt := config.OrgCert
	if strings.Contains(orgCrt, "goes here") == false {
		err = ioutil.WriteFile("/etc/chef/org.crt", []byte(orgCrt), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to write client config:%s\n", err)
		}
	}
	valPEM := config.ValidationPEM
	if strings.Contains(valPEM, "goes here") == false {
		err = ioutil.WriteFile("/etc/chef/validation.pem", []byte(valPEM), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to write client config:%s\n", err)
		}
	}

	// TODO: Check to see if chef is installed

	// Download chef client
	file, err := Download(GetChefURL(), ".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading file: %s\n", err)
		os.Exit(1)
	}

	var (
		pkgPath       string
		dmgMountPoint string
	)

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

	// TODO: This is very opinionated. Should it be removed or
	// controlled with a config option? Clayton
	// Set the firstboot tag to ensure the firstboot runlist is used.
	err = ioutil.WriteFile("/etc/chef/firstboot", []byte(""), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create firstboot file:%s\n", err)
	}

	_, err = RunChef()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Chef failed to complete: %s\n", err)
		os.Exit(1)
	}
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
