package client

import (
	"runtime"
	"testing"

	"github.com/clburlison/bakeit/src/config"
)

var darwinTestDynamic = "https://www.chef.io/chef/download/" +
	"?p=mac_os_x&pv=10.13&m=amd64&v=latest&prerelease=false"
var darwinTestStatic = "https://packages.chef.io/files/stable/" +
	"chef/13.6.4/mac_os_x/10.13/chef-13.6.4-1.dmg"
var windowsTestDynamic = "https://www.chef.io/chef/download/" +
	"?p=windows&pv=&m=i386&v=13.6.0&prerelease=false"
var windowsTestStatic = "https://packages.chef.io/files/stable/" +
	"chef/13.6.4/windows/2016/chef-client-13.6.4-1-x86.msi"

func TestGetChefURL(t *testing.T) {
	config.Verbose = false
	config.ChefClientVersion = "latest"
	switch os := runtime.GOOS; os {
	case "darwin":
		out := GetChefURL()
		if out != darwinTestDynamic {
			t.Errorf("GetChefURL: Have %s; want %s", out, darwinTestDynamic)
		}
		config.ChefClientURL["darwin"] = darwinTestStatic
		out = GetChefURL()
		if out != darwinTestStatic {
			t.Errorf("GetChefURL: Have %s; want %s", out, darwinTestStatic)
		}
	case "linux":
		// TODO: Write tests for linux. Need to figure out Debian vs FreeBSD vs ubuntu
	case "windows":
		out := GetChefURL()
		if out != windowsTestDynamic {
			t.Errorf("GetChefURL: Have %s; want %s", out, windowsTestDynamic)
		}
		config.ChefClientURL["windows"] = windowsTestStatic
		out = GetChefURL()
		if out != windowsTestStatic {
			t.Errorf("GetChefURL: Have %s; want %s", out, windowsTestStatic)
		}
	}
}
