package serial

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

// GetSerialNumber - Return the current serial number for the node or blank if
// unable to determine a serial number.
func GetSerialNumber() string {
	cmd := exec.Command("/usr/sbin/ioreg", "-c", "IOPlatformExpertDevice", "-d", "2")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("GetSerialNumber:", err)
	}
	r := regexp.MustCompile(`(?m)(?:IOPlatformSerialNumber" = ")([0-9A-Za-z]*)`)
	m := r.FindStringSubmatch(out.String())
	// The regex above should only give two capture groups when it works
	if len(m) == 2 {
		return m[1]
	}
	return ""
}
