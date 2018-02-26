package download

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/clburlison/bakeit/src/config"
)

type downloadList struct {
	url  string
	hash string
}

var urlID = "dcae73fa22fe9706a784a5b76d77dc3b"
var baseURL = fmt.Sprintf("https://gist.githubusercontent.com/clburlison"+
	"/%s/raw/", urlID)

var testDownloadList = []downloadList{
	{baseURL + "ab56d833409e921033f53fe45431d9c22030766a/hexdump",
		"14fd84a34d6dd237b8ec0d4c2caf3a738e9a39efcf0af9b7a85890676e0f452e"},
	{baseURL + "ab56d833409e921033f53fe45431d9c22030766a/uuid",
		"45705afc1b899a6f97fc12a255a796cdd969b171e9a9bc948b558ab1ba324ebe"},
}

func TestDownload(t *testing.T) {
	config.Verbose = true
	for i, test := range testDownloadList {
		file, err := ioutil.TempFile(os.TempDir(), "gotest_")
		defer os.Remove(file.Name())
		if err != nil {
			t.Errorf("#%d: Unable to create temp file %s\n", i, err)
		}
		_, err = Download(test.url, file.Name())
		if err != nil {
			t.Errorf("#%d: Unable to download '%s' file\n", i, path.Base(test.url))
		}
		status, err := CheckHash(file.Name(), test.hash)
		if err != nil {
			t.Errorf("#%d: Download(%s)=%t; want %s", i, path.Base(test.url), status, test.hash)
		}
	}
}
