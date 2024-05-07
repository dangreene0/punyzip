package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of punyzip.",
	Long:  `Print the version number of punyzip.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("punyzip tiny archive extractor v0.3")
	},
}
