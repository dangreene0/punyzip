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
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		msg, err := extract.DetectFormat(args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		fmt.Println(msg)

	},
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
