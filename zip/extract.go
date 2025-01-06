package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gonzabosio/vgstack-cli/ops"
)

func Download(templateRepoUrl, destPath string) error {
	resp, err := http.Get(templateRepoUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// store the downloaded .zip in new file
	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// copy the downloaded content to the .zip file
	_, err = io.Copy(out, resp.Body)
	return err
}

var ignoreDockerfiles = []string{"backend/Dockerfile", "backend/.dockerignore", "frontend/.dockerignore", "frontend/Dockerfile", "docker-compose.yml"}

func ExtractTemplateFolder(zipFilePath, destDir string, noDocker bool) error {
	zipReader, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer zipReader.Close()
	// loop zip files
	if noDocker {
		for _, file := range zipReader.File {
			if strings.HasPrefix(file.Name, "vgstack-temp-master/templates/") {
				// relative path by removing base
				relPath := file.Name[len("vgstack-temp-master/templates/"):]
				destPath := filepath.Join(destDir, relPath)
				if !ops.ShouldIgnore(relPath, ignoreDockerfiles) {
					// if the entry is a directory
					if file.FileInfo().IsDir() {
						if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
							return fmt.Errorf("extraction error: failed to create directory %v: %w", destPath, err)
						}
					} else {
						ops.CopyFile(file, destPath)
					}
				}
			}
		}
	} else {
		for _, file := range zipReader.File {
			if strings.HasPrefix(file.Name, "vgstack-temp-master/templates/") {
				// relative path by removing base
				relPath := file.Name[len("vgstack-temp-master/templates/"):]
				destPath := filepath.Join(destDir, relPath)
				// if the entry is a directory
				if file.FileInfo().IsDir() {
					if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
						return fmt.Errorf("extraction error: failed to create directory %v: %w", destPath, err)
					}
				} else {
					ops.CopyFile(file, destPath)
				}
			}
		}
	}
	return nil
}
