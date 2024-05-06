package extract

import (
	"compress/gzip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ExtractGz(archive string) error {
	fileName, f := strings.CutSuffix(filepath.Base(archive), ".gz")
	if !f {
		return errors.New(".gz suffix not found, potentially not real gzip file")
	}

	gzipStream, err := os.Open(archive)
	if err != nil {
		return err
	}
	defer gzipStream.Close()

	gzipReader, err := gzip.NewReader(gzipStream)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, gzipReader); err != nil {
		return err
	}

	return nil
}
