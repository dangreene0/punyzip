package extract

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
}

// make this take a pointer of the file
// :shrug:
func detectFormat() string {
	return "Detected was format X."
}

var ExtractCmd = &cobra.Command{
	Use:   "extract",
	Short: "extracts files.",
	Long:  `extracts a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(detectFormat())
	},
}
