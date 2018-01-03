package client

import (
	"bytes"
	"os/exec"
	"strings"
)

// MacInfoObject - Hold mac specific data from sw_vers
type MacInfoObject struct {
	ProductName    string
	ProductVersion string
	BuildVersion   string
}

// GetMacInfo - return facts about current node
func GetMacInfo() (*MacInfoObject, error) {
	cmd := exec.Command("/usr/bin/sw_vers")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	data := strings.Split(out.String(), "\n")
	productName := strings.TrimSpace(strings.Split(data[0], ":")[1])
	productVersion := strings.TrimSpace(strings.Split(data[1], ":")[1])
	buildVersion := strings.TrimSpace(strings.Split(data[2], ":")[1])
	mio := &MacInfoObject{productName, productVersion, buildVersion}
	return mio, nil
}
