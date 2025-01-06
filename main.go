package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/gonzabosio/vgstack-cli/zip"
)

const (
	tempRepoUrl = "https://github.com/gonzabosio/vgstack-temp/archive/refs/heads/master.zip"
	zipFilePath = "temp.zip"
	tempDstDir  = "./"
	backendSrc  = "./templates/backend"
	frontendSrc = "./templates/frontend"
)

func main() {
	noDocker := flag.Bool("nodocker", false, "Use -nodocker if you want to disable docker files generation")
	flag.Parse()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := zip.Download(tempRepoUrl, zipFilePath)
		if err != nil {
			fmt.Println("Error downloading zip file:", err)
			return
		}

		err = zip.ExtractTemplateFolder(zipFilePath, tempDstDir, *noDocker)
		if err != nil {
			fmt.Println("Error extracting templates:", err)
			return
		}
	}()
	wg.Wait()

	if err := os.Remove("temp.zip"); err != nil {
		fmt.Println("Error removing template zip file: ", err)
		return
	}

	fmt.Printf("Project was created successfully! 🎉\nJust one more thing!\nGo to the 'README.md' file and use the commands specified")
}
