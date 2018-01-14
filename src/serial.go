package client

import "regexp"

// CleanNodeNameChars will remove unsuppported characters.
// Chef only supports the following for node name:
// A-Z, a-z, 0-9, _, -, or .
func CleanNodeNameChars(serial string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9._-]+")
	if err != nil {
		return "", err
	}
	processedString := reg.ReplaceAllString(serial, "")
	return processedString, nil
}
