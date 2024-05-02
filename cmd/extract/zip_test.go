package extract

import (
	"os"
	"testing"
)

func TestExtractZip(t *testing.T) {

	testCases := []struct {
		input    string
		expected []string
	}{
		{"examples/archive.zip", []string{"baby_slime.png", "go_piece.jpg", "gopher.png"}},
		{"examples/archive_directories.zip", []string{"root/baby_slime.png", "root/subdirectory/go_piece.jpg", "gopher.png"}},
	}

	for _, testCase := range testCases {
		err := ExtractZip("../../" + testCase.input)
		if err != nil {
			t.Errorf(err.Error())
		}
		for _, item := range testCase.expected {
			_, err = os.Stat("../../" + item)
			if err != nil {
				t.Errorf("Expected %q, got none", item)
			}
		}

	}
}
