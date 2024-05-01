package extract

import (
	"testing"
)

func TestDetectFormat(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"examples/archive.tar.gz", "Extracted Tar File"},
		{"examples/archive.7z", "Extracted 7Zip File"},
		{"examples/archive.rar", "Extracted RAR File"},
		{"examples/archive.bz2", ""},
		{"examples/archive.zip", "Extracted Zip File"},
		{"examples/archive.zip.exe", ""},
		{"examples/archive.exe.zip", "Extracted Zip File"},
		{"examples/archive.tar.zip", "Extracted Zip File"},
		{"examples/zip.zip", "Extracted Zip File"},
		{"examples/zip.7z", "Extracted 7Zip File"},
		{"examples/archive.jpg", ""},
		{"examples/archive.png", ""},
		{"examples/archive", ""},
	}

	for _, testCase := range testCases {
		got, _ := DetectFormat(testCase.input)

		if testCase.expected != got {
			t.Errorf("Expected %q, got %q", testCase.expected, got)
		}
	}
}
