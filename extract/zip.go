package extract

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ExtractZip(archive string) error {
	r, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if err = extractZFile(f); err != nil {
			return err
		}
	}
	return nil
}

func extractZFile(file *zip.File) error {

	if file.FileInfo().IsDir() {
		// Create the directory
		fmt.Printf("creating directory %s\n", file.Name)
		err := os.MkdirAll(file.Name, os.ModePerm)
		if err != nil {
			return err
		}
		return nil
	}

	// Create the parent directory if it doesn't exist
	err := os.MkdirAll(filepath.Dir(file.Name), os.ModePerm)
	if err != nil {
		return err
	}

	srcFile, err := file.Open()
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// set it to extract back in it's folder
	// Create an empty destination file
	dstFile, err := os.OpenFile(file.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
