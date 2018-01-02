package client

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/clburlison/bakeit/client/config"
)

var chefBaseURL = "https://www.chef.io/chef/download/"

func getPlatformInfo() (ver string, plat string) {
	var osVers string
	switch os := runtime.GOOS; os {
	case "darwin":
		osVers = chop(GetMacInfo().ProductVersion, 2)
		return osVers, "mac_os_x"
	case "linux":
		fmt.Println("Linux")
		return osVers, "linux"
	case "windows":
		return "", "windows"
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	return "UNKNOWN", "UNKNOWN"
}

// getChefWebURL will build a url to download directly from
// chef.io for the specific platform.
func getChefWebURL() string {
	ver, plat := getPlatformInfo()
	arch := runtime.GOARCH
	if arch == "386" {
		arch = "i386"
	}
	s := []string{chefBaseURL, "?p=", plat,
		"&pv=", ver, "&m=",
		arch, "&v=", config.ChefClientVersion,
		"&prerelease=", config.ChefClientPreRelease}
	return strings.Join(s, "")
}

// GetChefURL will return a platform specific URL from config or a built URL from getChefWebURL
func GetChefURL() string {
	var chefURL string
	configURL := config.ChefClientURL[runtime.GOOS]
	if configURL != "" {
		chefURL = configURL
	} else {
		chefURL = getChefWebURL()
	}
	return chefURL
}

type macInfoObject struct {
	ProductName    string
	ProductVersion string
	BuildVersion   string
}

func GetMacInfo() *macInfoObject {
	cmd := exec.Command("/usr/bin/sw_vers")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("GetMacInfo:", err)
	}
	data := strings.Split(out.String(), "\n")
	productName := strings.TrimSpace(strings.Split(data[0], ":")[1])
	productVersion := strings.TrimSpace(strings.Split(data[1], ":")[1])
	buildVersion := strings.TrimSpace(strings.Split(data[2], ":")[1])
	mio := &macInfoObject{productName, productVersion, buildVersion}
	return mio
}

func chop(s string, i int) string {
	return s[0 : len(s)-i]
}
