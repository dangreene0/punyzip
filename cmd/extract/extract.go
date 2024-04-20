package extract

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
}

// make this take a pointer of the file
// :shrug:
func detectFormat(archive string) (string, error) {
	switch {
	case strings.HasSuffix(archive, ".zip"):
		return "zip", nil
	case strings.HasSuffix(archive, ".tar.gz"):
		return "tar", nil
	case strings.HasSuffix(archive, ".7z"):
		return "7z", nil
	case strings.HasSuffix(archive, ".rar"):
		return "rar", nil
	}
	return "", errors.New("unsupported archive format")
}

// consider having no commands and having it autodetect everything for the user
var ExtractCmd = &cobra.Command{
	Use:   "x",
	Short: "extracts files.",
	Long:  `extracts a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := extractZip(args[0])
		if err != nil {
			panic(err)
		}
		fmt.Println(msg)
	},
}
