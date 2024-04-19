package extract

import (
	"fmt"

	"github.com/spf13/cobra"
)

func extractZip() string {
	return "Extracted zip file"
}

var zipCmd = &cobra.Command{
	Use:   "zip",
	Short: "extracts zip files.",
	Long:  `extracts zip files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(extractZip())
	},
}

func init() {
	ExtractCmd.AddCommand(zipCmd)
}
