package filetype

import "testing"

type TestFileType struct {
	path string
	ext  string
	mime string
}

var testFileType = []TestFileType{
	// TODO: Unclear why the pkg test fails randomly.
	// {"./testdata/echo.pkg", "xar", "application/x-xar"},
	{"./testdata/echo.dmg", "zlib", "application/zlib"},
}

func TestMatch(t *testing.T) {
	for i, test := range testFileType {
		ext, mime := Match(test.path)
		if ext != test.ext {
			t.Errorf("#%d: Match(%s)=%s; want %s\n", i, test.path, ext, test.ext)
		}
		if mime != test.mime {
			t.Errorf("#%d: Match(%s)=%s; want %s\n", i, test.path, mime, test.mime)
		}
	}
}
