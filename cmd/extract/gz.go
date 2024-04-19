package extract

import (
	"fmt"

	"github.com/spf13/cobra"
)

func extractGZ() string {
	return "Extracted gz file"
}

var gZCmd = &cobra.Command{
	Use:   "gz",
	Short: "extracts gz files.",
	Long:  `extracts gz files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(extractGZ())
	},
}

func init() {
	ExtractCmd.AddCommand(gZCmd)
}
