package serial

// GetSerialNumber - Return the current serial number for the node or blank if
// unable to determine a serial number.
func GetSerialNumber() string {
	// dmidecode -t system | grep Serial
	return ""
}
