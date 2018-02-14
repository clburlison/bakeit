package node

import (
	"bytes"
	"os/exec"
	"strings"
)

// GetNodeInfo - return facts about current node
func GetNodeInfo() (*InfoObject, error) {
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
	mio := &InfoObject{productName, productVersion, buildVersion}
	return mio, nil
}
