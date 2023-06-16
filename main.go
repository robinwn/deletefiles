package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Check if the directory prefix is provided as a command line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: deletefiles <root_dir>")
		os.Exit(1)
	}

	// Get the directory prefix from the command line arguments
	directoryPrefix := os.Args[1]

	// Read file names from stdin
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Get the relative file name
		relativeFileName := scanner.Text()

		// Construct the absolute file path
		absoluteFilePath := filepath.Join(directoryPrefix, relativeFileName)

		// Delete the file
		err := os.Remove(absoluteFilePath)
		if err != nil {
			fmt.Printf("Error deleting file '%s': %v\n", absoluteFilePath, err)
		} else {
			fmt.Printf("Deleted file: %s\n", absoluteFilePath)
		}
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
