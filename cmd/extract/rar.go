package extract

import (
	"fmt"

	"github.com/spf13/cobra"
)

func extractRar() string {
	return "Extracted rar file"
}

var rarCmd = &cobra.Command{
	Use:   "rar",
	Short: "extracts rar files.",
	Long:  `extracts rar files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(extractRar())
	},
}

func init() {
	ExtractCmd.AddCommand(rarCmd)
}
