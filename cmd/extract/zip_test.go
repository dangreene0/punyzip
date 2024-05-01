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
	}

	for _, testCase := range testCases {
		err := ExtractZip(testCase.input)
		if err != nil {
			t.Errorf(err.Error())
		}
		_, err = os.Stat("example/" + testCase.expected[0])

		if err != nil {
			t.Errorf("Expected %q, got none", "example/"+testCase.expected[0])
		}
	}
}
