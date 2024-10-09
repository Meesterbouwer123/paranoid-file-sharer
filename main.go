package main

import (
	"fmt"
	"log"
	"os"
	"paranoid-file-sharer/backend"
	fileencryption "paranoid-file-sharer/utils/file_encryption"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting up...")
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init server-side encryption
	fileencryption.Init()

	// make uploads directory if it doesn't exist yet
	if _, err := os.Stat("uploads/"); os.IsNotExist(err) {
		if err := os.Mkdir("uploads/", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	// launch site
	fmt.Println("Launching site")
	backend.StartBackend()
}
