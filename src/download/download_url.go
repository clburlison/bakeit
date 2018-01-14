package download

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/clburlison/bakeit/src/config"
	"github.com/clburlison/bakeit/src/node"
)

var chefBaseURL = "https://www.chef.io/chef/download/"

func getPlatformInfo() (ver string, plat string) {
	var osVers string
	switch os := runtime.GOOS; os {
	case "darwin":
		info, err := node.GetNodeInfo()
		if err != nil {
			fmt.Printf("Unable to obtain macOS version info: %s\n", err)
		}
		osVers = node.Chop(info.ProductVersion, 2)
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
