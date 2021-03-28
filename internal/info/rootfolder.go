package info

import (
	"os"
	"path/filepath"
)

// RootFolder looks for the root folder by scanning from the current directory and up
// for the first occurrence of go.mod, if it does not find the module returns empty string.
func RootFolder() string {
	basePath, err := os.Getwd()
	if err != nil {
		return ""
	}

	targetPath := basePath
	for {
		_, err := filepath.Rel(basePath, targetPath)
		if err != nil {
			return ""
		}

		abso, err := filepath.Abs(targetPath)
		if err != nil {
			return ""
		}

		path := filepath.Join(targetPath, "go.mod")
		_, err = os.Stat(path)

		if err != nil && abso == "/" {
			break
		}

		if err != nil {
			targetPath += "/.."

			continue
		}

		return abso
	}

	return ""
}
