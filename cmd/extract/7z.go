package extract

import (
	"fmt"

	"github.com/bodgit/sevenzip"
	"github.com/spf13/cobra"
)

func extractFile(file *sevenzip.File) error {
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	return nil
}

func extractSevenZip(archive string) error {
	r, err := sevenzip.OpenReader(archive)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if err = extractFile(f); err != nil {
			return err
		}
	}

	return nil
}

var sevenZCmd = &cobra.Command{
	Use:   "7z",
	Short: "extracts 7z files.",
	Long:  `extracts 7z files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(extractSevenZip(args[0]))
	},
}

func init() {
	ExtractCmd.AddCommand(sevenZCmd)
}
