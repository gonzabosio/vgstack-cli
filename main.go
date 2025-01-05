package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	backendSrc  = "./templates/backend"
	frontendSrc = "./templates/frontend"
	makefileSrc = "./templates/Makefile"
	composeSrc  = "./templates/docker-compose.yml"
)

func main() {
	frontendName := ""
	backendName := ""
	flag.StringVar(&backendName, "b", "", "Type the name of your backend folder. e.g: -b 'back'")
	flag.StringVar(&frontendName, "f", "", "Type the name for your frontend foldedr. e.g: -f 'front'")
	noDocker := flag.Bool("nodocker", false, "Use nodocker if you want to disable docker files generation")
	flag.Parse()

	if backendName != "" {
		if err := copyDir(backendSrc, "./"+backendName, *noDocker); err != nil {
			log.Fatalf("failed to copy backend files: %v", err)
		}
	} else {
		backendName = "backend"
		if err := copyDir(backendSrc, "./"+backendName, *noDocker); err != nil {
			log.Fatalf("failed to copy backend files: %v", err)
		}
	}
	if frontendName != "" {
		if err := copyDir(frontendSrc, "./"+frontendName, *noDocker); err != nil {
			log.Fatalf("failed to copy frontend files: %v", err)
		}
	} else {
		frontendName = "frontend"
		if err := copyDir(frontendSrc, "./"+frontendName, *noDocker); err != nil {
			log.Fatalf("failed to copy frontend files: %v", err)
		}
	}
	if !*noDocker {
		src, err := os.Open(composeSrc)
		if err != nil {
			log.Fatalf("failed to open the source docker-compose")
		}
		dst, err := os.Create("./docker-compose.yml")
		if err != nil {
			log.Fatalf("failed to create docker-compose file")
		}
		_, err = io.Copy(dst, src)
		if err != nil {
			log.Fatalf("failed to copy docker-compose file")
		}
	}

	src, err := os.Open(makefileSrc)
	if err != nil {
		log.Fatalf("failed to open the source makefile")
	}
	dst, err := os.Create("./Makefile")
	if err != nil {
		log.Fatalf("failed to create makefile")
	}
	_, err = io.Copy(dst, src)
	if err != nil {
		log.Fatalf("failed to copy makefile")
	}

	fmt.Printf("Project was created successfully! 🎉\nJust one more thing!\nGo to the 'README.md' file and use the commands specified")
}
