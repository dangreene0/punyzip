package cmd

import (
	"fmt"
	"os"

	"github.com/dangreene0/punyzip/cmd/extract"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "punyzip",
	Short: "A tiny archive extractor.",
	Long:  "A tiny archive extractor that supports zip, 7z, rar, tar and tar.gz files.",
	Run: func(cmd *cobra.Command, args []string) {
		panic("Not implemented yet.")
	},
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(extract.ExtractCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
