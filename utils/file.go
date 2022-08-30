package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func DownloadFile(url string, path string) error {
	var dir = filepath.Dir(path)
	if !IsExist(dir) {
		os.MkdirAll(dir, os.ModePerm)
	}
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
