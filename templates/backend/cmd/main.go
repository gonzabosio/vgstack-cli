package main

import (
	"log"
	backend "vgstack-cli/templates/backend/api"

	"github.com/joho/godotenv"
)

func main() {
	// UNCOMMENT .ENV FROM THE GITIGNORE FILE

	if err := godotenv.Load("backend/.env"); err != nil {
		log.Fatalf("environment variables loading error: %v", err)
	}
	if err := backend.StartBackendServer(); err != nil {
		log.Fatalf("Unable to start backend server: %v", err)
	}
}
