package extract

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ExtractZip(file string) error {

	archive, err := zip.OpenReader(file)
	if err != nil {
		return err
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
				return err
			}
			continue
		}

		// Create the parent directory if it doesn't exist
		err := os.MkdirAll(filepath.Dir(f.Name), os.ModePerm)
		if err != nil {
			return err
		}

		// // Create an empty destination file
		dstFile, err := os.OpenFile(f.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer dstFile.Close()

		// Open the file in the zip and copy its contents to the destination file
		srcFile, err := f.Open()
		if err != nil {
			return err
		}
		defer srcFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}
	}
	return nil
}
