package main

import (
	"bufio"
	"chatserver/server"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func printBanner(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening the file: %s\n", err)
		return
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Print the ASCII art on the command line
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading the file: %s\n", err)
		return
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading environment variables: ", err)
	}

	printBanner("art/banner.txt")
	url := os.Getenv("ADDRESS")
	log.Println("Starting server on: ", url)
	server.InitServer(url)
}
