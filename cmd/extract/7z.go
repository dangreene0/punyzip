package extract

import (
	"fmt"

	"github.com/spf13/cobra"
)

func extract7Z() string {
	return "Extracted 7z file"
}

var sevenZCmd = &cobra.Command{
	Use:   "7z",
	Short: "extracts 7z files.",
	Long:  `extracts 7z files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(extract7Z())
	},
}

func init() {
	ExtractCmd.AddCommand(sevenZCmd)
}
