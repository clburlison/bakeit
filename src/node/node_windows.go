package node

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// GetNodeInfo empty func to allow CI to complete
func GetNodeInfo() (*InfoObject, error) {
	return nil, nil
}

// GetWin32OS exports win32_operatingsystem powershell class
func GetWin32OS() (Win32OS, error) {
	cmd := exec.Command("powershell", "gwmi", "Win32_OperatingSystem", "|", "ConvertTo-Json")

	// cmd.Stderr = os.Stderr
	o, err := cmd.Output()

	if err != nil {
		return Win32OS{}, fmt.Errorf("exec gwmi Win32_OperatingSystem: %s", err)
	}

	var j Win32OS

	if err := json.Unmarshal(o, &j); err != nil {
		return Win32OS{}, fmt.Errorf("failed unmarshalling Win32_OperatingSystem: %s", err)
	}

	return j, nil
}

// Win32OS structure
type Win32OS struct {
	Caption                string `json:"Caption"` //os version
	TotalVirtualMemorySize int    `json:"TotalVirtualMemorySize"`
	TotalVisibleMemorySize int    `json:"TotalVisibleMemorySize"`
}
