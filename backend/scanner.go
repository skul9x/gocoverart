package backend

import (
	"os"
	"path/filepath"
	"strings"
)

// ScanMusicFiles scans a directory for mp3 and flac files
func ScanMusicFiles(dirPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))
			if ext == ".mp3" || ext == ".flac" {
				files = append(files, path)
			}
		}
		return nil
	})

	return files, err
}
