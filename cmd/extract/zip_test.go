package extract

import "testing"

func TestExtractZip(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"examples/archive.zip", "zip"},
		{"examples/archive.tar.gz", "tar"},
		{"examples/archive.7z", "7z"},
		{"examples/archive.rar", "rar"},
		{"examples/archive.bz2", ""},
		{"examples/archive.tar.gz.zip", ""},
		{"examples/archive.jpg", ""},
		{"examples/archive.png", ""},
		{"examples/archive", ""},
	}

	for _, testCase := range testCases {
		got, _ := extractZip(testCase.input)

		if testCase.expected != got {
			t.Errorf("Expected %q, got %q", testCase.expected, got)
		}
	}
}
