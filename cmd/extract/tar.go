package extract

import (
	"fmt"

	"github.com/spf13/cobra"
)

func extractTar() string {
	return "Extracted tar file"
}

var tarCmd = &cobra.Command{
	Use:   "tar",
	Short: "extracts tar files.",
	Long:  `extracts tar files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(extractTar())
	},
}

func init() {
	ExtractCmd.AddCommand(tarCmd)
}
