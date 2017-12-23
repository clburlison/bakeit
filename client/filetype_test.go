package client

import "testing"

type Test struct {
	path string
	ext  string
	mime string
}

var tests = []Test{
	{"./testdata/echo.pkg", "xar", "application/x-xar"},
	{"./testdata/echo.dmg", "zlib", "application/zlib"},
}

func TestMatch(t *testing.T) {
	for i, test := range tests {
		ext, mime := Match(test.path)
		if ext != test.ext {
			t.Errorf("#%d: Match(%s)=%s; want %s", i, test.path, ext, test.ext)
		}
		if mime != test.mime {
			t.Errorf("#%d: Match(%s)=%s; want %s", i, test.path, mime, test.mime)
		}
	}
}
