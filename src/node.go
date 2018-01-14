package client

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/clburlison/bakeit/src/config"
)

// NodeInfoObject - Hold node specific data
type NodeInfoObject struct {
	ProductName    string
	ProductVersion string
	BuildVersion   string
}

// Chop - Easily delete trailing characters from a string
func Chop(s string, i int) string {
	return s[0 : len(s)-i]
}

var (
	chefRootPath = map[string]string{
		"darwin":  "/etc/chef",
		"windows": "C:\\chef",
		"linux":   "",
	}

	chefClientFile = map[string]string{
		"darwin":  "/etc/chef/client.rb",
		"windows": "C:\\chef\\client.rb",
		"linux":   "",
	}

	chefOrgCertPath = map[string]string{
		"darwin":  "/etc/chef/org.crt",
		"windows": "C:\\chef\\org.crt",
		"linux":   "",
	}
)

// ChefFiles writes required config files to disk
func ChefFiles(clientConfig string) error {
	// Remove old chef files before running
	if config.Force {
		clientCert := config.ChefClientCertPath[runtime.GOOS]
		if _, err := os.Stat(clientCert); os.IsExist(err) {
			os.Remove(clientCert)
		}
		ohaiPlugins := config.ChefClientOhaiDirectory[runtime.GOOS]
		if _, err := os.Stat(ohaiPlugins); os.IsExist(err) {
			os.Remove(ohaiPlugins)
		}
	}

	// Write chef config files
	if _, err := os.Stat(chefRootPath[runtime.GOOS]); os.IsNotExist(err) {
		os.MkdirAll(chefRootPath[runtime.GOOS], os.ModePerm)
	}
	err := ioutil.WriteFile(chefClientFile[runtime.GOOS], []byte(clientConfig), 0644)
	if err != nil {
		return fmt.Errorf("Unable to write client config:%s", err)
	}
	json := config.ChefClientRunListJSON["darwin"]
	err = ioutil.WriteFile(config.ChefClientJSONAttribs[runtime.GOOS], []byte(json), 0644)
	if err != nil {
		return fmt.Errorf("Error formating JSON run list:%s", err)
	}
	orgCrt := config.OrgCert
	if strings.Contains(orgCrt, "goes here") == false {
		err = ioutil.WriteFile(chefOrgCertPath[runtime.GOOS], []byte(orgCrt), 0644)
		if err != nil {
			return fmt.Errorf("Unable to write client config:%s", err)
		}
	}
	valPEM := config.ValidationPEM
	if strings.Contains(valPEM, "goes here") == false {
		err = ioutil.WriteFile(config.ChefClientValidationKey[runtime.GOOS], []byte(valPEM), 0644)
		if err != nil {
			return fmt.Errorf("Unable to write client config:%s", err)
		}
	}
	return nil
}
