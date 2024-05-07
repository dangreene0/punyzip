package extract

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var ExtractCmd = &cobra.Command{
	Use:   "x",
	Short: "extracts files.",
	Long:  `extracts a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := DetectFormat(args[0])
		if err != nil {
			panic(err)
		}
		fmt.Println(msg)
	},
}

// ACTUALLY CHECK FILE HEADER!
func DetectFormat(archive string) (string, error) {
	switch {
	case strings.HasSuffix(archive, ".zip"):
		return "Extracted Zip File", ExtractZip(archive)
	case strings.HasSuffix(archive, "tar.gz"):
		err := ExtractGz(archive)
		if err != nil {
			return "Failure", err
		}
		tar, _ := strings.CutSuffix(filepath.Base(archive), ".gz")
		return "Extracted Compressed Tarball File", ExtractTar(tar)
	case strings.HasSuffix(archive, ".gz"):
		return "Extracted Gzip File", ExtractGz(archive)
	case strings.HasSuffix(archive, ".tar"):
		return "Extracted Tar File", ExtractTar(archive)
	case strings.HasSuffix(archive, ".7z"):
		return "Extracted 7Zip File", ExtractSevenZip(archive)
	case strings.HasSuffix(archive, ".rar"):
		return "Extracted RAR File", nil
	}
	return "", errors.New("unsupported archive format")
}
