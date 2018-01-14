package serial

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// GetSerialNumber - Return the current serial number for the node
func GetSerialNumber() string {
	cmd := exec.Command("wmic", "bios", "get", "serialnumber")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("GetSerialNumber:", err)
	}
	serial := strings.TrimPrefix(out.String(), "SerialNumber")
	serial = strings.TrimSpace(serial)
	return serial
}
