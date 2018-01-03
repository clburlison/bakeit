package client

import (
	"io/ioutil"
	"regexp"
	"strings"
	"testing"
)

var macSerial = "C00BB1CDDEEF"
var windowsSerial = "VMware-56 4d c6 6a 64 96 72 56-67 d0 81 f6 80 6b 53 84"

func TestGetSerialNumber(t *testing.T) {
	// Mac Test
	a, err := ioutil.ReadFile("testdata/serial_darwin")
	if err != nil {
		t.Errorf("'serial_darwin' is missing!\n")
	}
	darwinSerialDump := string(a)
	r := regexp.MustCompile(`(?m)(?:IOPlatformSerialNumber" = ")([0-9A-Za-z]*)`)
	m := r.FindStringSubmatch(darwinSerialDump)
	if m[1] != macSerial {
		t.Errorf("GetSerialNumber: Have %s; want %s", m[1], macSerial)
	}

	// Windows Test
	b, err := ioutil.ReadFile("testdata/serial_windows")
	if err != nil {
		t.Errorf("'serial_windows' is missing!\n")
	}
	windowsSerialDump := string(b)
	winSerial := strings.TrimPrefix(windowsSerialDump, "SerialNumber")
	winSerial = strings.TrimSpace(winSerial)
	if winSerial != windowsSerial {
		t.Errorf("GetSerialNumber: Have %s; want %s", winSerial, windowsSerial)
	}
}
