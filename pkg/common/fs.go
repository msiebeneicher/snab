package common

import (
	"os"
)

// IsDirectory determines if a file represented
// by `path` is a directory or not
func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

// IsFile determines if a file represented
// by `path` is a file or not
func IsFile(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, err
	}
	return !fileInfo.IsDir(), err
}
