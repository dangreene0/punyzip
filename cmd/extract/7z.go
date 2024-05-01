package extract

import (
	"github.com/bodgit/sevenzip"
)

func extractFile(file *sevenzip.File) error {
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	return nil
}

func ExtractSevenZip(archive string) error {
	r, err := sevenzip.OpenReader(archive)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if err = extractFile(f); err != nil {
			return err
		}
	}

	return nil
}
