package extract

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func extractZip(file string) (string, error) {
	format, _ := detectFormat(file)
	if format != "zip" {
		return "", errors.New(file + " is not correct Archive Format.")
	}

	archive, err := zip.OpenReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	// Extract the files from the zip
	for _, f := range archive.File {

		// Print the file path
		fmt.Println("extracting file ", f.Name)

		// Check if the file is a directory
		if f.FileInfo().IsDir() {
			// Create the directory
			fmt.Println("creating directory...")
			err := os.MkdirAll(f.Name, os.ModePerm)
			if err != nil {
				panic(err)
			}
			continue
		}

		// Create the parent directory if it doesn't exist
		err := os.MkdirAll(filepath.Dir(f.Name), os.ModePerm)
		if err != nil {
			panic(err)
		}

		// // Create an empty destination file
		dstFile, err := os.OpenFile(f.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}
		defer dstFile.Close()

		// Open the file in the zip and copy its contents to the destination file
		srcFile, err := f.Open()
		if err != nil {
			panic(err)
		}
		defer srcFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			panic(err)
		}
	}
	return "Extracted zip file", nil
}

var zipCmd = &cobra.Command{
	Use:   "zip",
	Short: "extracts zip files.",
	Long:  `extracts zip files.`,
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := extractZip(args[0])
		if err != nil {
			panic(err)
		}
		fmt.Println(msg)
	},
}

func init() {
	ExtractCmd.AddCommand(zipCmd)
}
