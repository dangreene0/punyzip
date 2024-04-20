package extract

import (
	"testing"
)

func TestDetectFormat(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"archive.tar.gz", "tar"},
		{"archive.7z", "7z"},
		{"archive.rar", "rar"},
		{"archive.bz2", ""},
		{"archive.zip", "zip"},
		{"archive.zip.exe", ""},
		{"archive.exe.zip", "zip"},
		{"archive.tar.zip", "zip"},
		{"zip.zip", "zip"},
		{"zip.7z", "7z"},
		{"archive.jpg", ""},
		{"archive.png", ""},
		{"archive", ""},
	}

	for _, testCase := range testCases {
		got, _ := detectFormat(testCase.input)

		if testCase.expected != got {
			t.Errorf("Expected %q, got %q", testCase.expected, got)
		}
	}
}
