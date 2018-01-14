package filetype

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"math"

	"gopkg.in/h2non/filetype.v1"
)

var xarType = filetype.NewType("xar", "application/x-xar")
var zlibType = filetype.NewType("zlib", "application/zlib")

// For .pkg files
func xarMatcher(buf []byte) bool {
	return len(buf) > 1 &&
		buf[0] == 0x78 && buf[1] == 0x61 && buf[2] == 0x72
}

// For .dmg files based off rfc6713
func zlibMatcher(buf []byte) bool {
	headBytes := readInt16([]byte{buf[0], buf[1]})
	return len(buf) > 1 &&
		buf[0] == 0x78 || buf[0] == 0x08 || buf[0] == 0x18 ||
		buf[0] == 0x28 || buf[0] == 0x38 || buf[0] == 0x48 ||
		buf[0] == 0x58 || buf[0] == 0x68 &&
		divisibleBy(headBytes, 31) == true
}

func readInt16(data []byte) (ret int16) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &ret)
	return
}

func divisibleBy(num int16, divisor int) bool {
	if math.Mod(float64(num), float64(divisor)) == 0 {
		return true
	}
	return false
}

// Match takes a local file path; returns file type, & mime
func Match(path string) (Extension string, MIME string) {
	filetype.AddMatcher(xarType, xarMatcher)
	filetype.AddMatcher(zlibType, zlibMatcher)

	buf, _ := ioutil.ReadFile(path)
	k, _ := filetype.Match(buf)

	return k.Extension, k.MIME.Value
}
