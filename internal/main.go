package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s fetch|format\n", os.Args[0])
		return
	}
	cmd := os.Args[1]
	switch cmd {
	case "fetch":
		if err := fetchRecords("/tmp/records.txt"); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to fetch records: %v\n", err)
			os.Exit(1)
		}
	case "format":
		if err := formatRecords("/tmp/records.txt", "/tmp/formatted.txt"); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to format records: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Invalid command: %s\n", cmd)
		os.Exit(1)
	}
}
