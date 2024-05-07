package extract

import (
	"archive/tar"
	"errors"
	"fmt"
	"io"
	"os"
)

func ExtractTar(archive string) error {
	tarStream, err := os.Open(archive)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(tarStream)
	fmt.Println(tarReader)
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(header.Name, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			outFile, err := os.Create(header.Name)
			if err != nil {
				return err
			}
			defer outFile.Close()

			// if _, err := io.Copy(outFile, tarReader); err != nil {
			// 	return err
			// }

		default:
			return errors.New(fmt.Sprintf(
				"ExtractTarGz: uknown type: %v in %s",
				header.Typeflag,
				header.Name))
		}
	}
	return nil
}
