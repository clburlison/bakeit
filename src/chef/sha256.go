package chef

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// CheckHash - Accept a local file path and expected sha256 hash. Returns bool
// if the hash matches the expected hash.
func CheckHash(file string, sha string) (bool, error) {
	f, err := os.Open(file)
	if err != nil {
		return false, fmt.Errorf("Unable to open file %s", err)
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		// fmt.("#%d: Unable to copy download to temp file %s\n", i, err)
		return false, err
	}
	shaHash := hex.EncodeToString(h.Sum(nil))
	if shaHash != sha {
		return false, fmt.Errorf("Downloaded file hash does not match config hash")
	}
	return true, nil
}
