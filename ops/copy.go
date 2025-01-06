package ops

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

// CopyFile copies a single file from source to destiny.
func CopyFile(src *zip.File, dst string) error {
	zipFile, err := src.Open()
	if err != nil {
		return fmt.Errorf("copy file error: failed to open file %v: %w", src.Name, err)
	}
	defer zipFile.Close()

	// Create the destination file
	destFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("copy file error: failed to create file %v: %w", dst, err)
	}
	defer destFile.Close()

	// Copy the content
	if _, err := io.Copy(destFile, zipFile); err != nil {
		return fmt.Errorf("copy file error: failed to copy file content to %v: %w", dst, err)
	}
	return nil
}

func ShouldIgnore(filename string, ignoreFiles []string) bool {
	for _, f := range ignoreFiles {
		fmt.Println("to ignore:", f, "have:", filename, "so:", f == filename)
		if filename == f {
			return true
		}
	}
	return false
}
