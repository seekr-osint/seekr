package seekrplugin

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func unpackFile(path, file string) (string, error) {
	tarball, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer tarball.Close()
	gzipReader, err := gzip.NewReader(tarball)
	if err != nil {
		return "", err
	}
	defer gzipReader.Close()

	tempDir, err := os.MkdirTemp("", "seekrBundle-")
	if err != nil {
		return "", err
	}

	tarReader := tar.NewReader(gzipReader)

	foundFile := false
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		if header.Name == file {
			foundFile = true
			break
		}
	}

	if !foundFile {
		log.Printf("File %s not found\n", file)
		return "", os.ErrNotExist
	}

	fileDir := filepath.Dir(filepath.Join(tempDir, file))
	err = os.MkdirAll(fileDir, 0755)
	if err != nil {
		return "", err
	}

	outFile, err := os.Create(filepath.Join(tempDir, file))
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, tarReader)
	if err != nil {
		return "", err
	}

	return tempDir, nil
}
