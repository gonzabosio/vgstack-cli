package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// copyFile copies a single file from source to destiny.
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	// preserve file permissions
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	return os.Chmod(dst, info.Mode())
}

// CopyDir copies a directory and its contents from source to destiny.
func copyDir(src, dst string) error {
	// get properties/description of the source directory
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !srcInfo.IsDir() {
		return fmt.Errorf("%s is not a directory", src)
	}

	err = os.MkdirAll(dst, srcInfo.Mode())
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		// fmt.Printf("- src: %v - dst: %v\n", srcPath, dstPath)
		if entry.IsDir() {
			// copy subdirs
			err = copyDir(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			// copy individual files
			err = copyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
