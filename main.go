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
)

func main() {
	frontendName := ""
	backendName := ""
	flag.StringVar(&backendName, "b", "", "Type the name of your backend folder. e.g: -b 'my_backend'")
	flag.StringVar(&frontendName, "f", "", "Type the name for your frontend. e.g: -f 'my_frontend'")
	flag.Parse()

	if backendName != "" {
		if err := copyDir(backendSrc, "./"+backendName); err != nil {
			log.Fatalf("failed to copy backend files: %v", err)
		}
	} else {
		backendName = "backend"
		if err := copyDir(backendSrc, "./"+backendName); err != nil {
			log.Fatalf("failed to copy backend files: %v", err)
		}
	}
	if frontendName != "" {
		if err := copyDir(frontendSrc, "./"+frontendName); err != nil {
			log.Fatalf("failed to copy frontend files: %v", err)
		}
	} else {
		frontendName = "frontend"
		if err := copyDir(frontendSrc, "./"+frontendName); err != nil {
			log.Fatalf("failed to copy frontend files: %v", err)
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
