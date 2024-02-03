package main

import (
	"testing"
	"os"
)

func TestGetCurrentDirectoryIsValidPath(t *testing.T) {
	// Get the current working directory
	currentDir, err := getCurrentDirectory()

	// Check for errors
	if err != nil {
		t.Fatalf("Error getting current working directory: %v", err)
	}

	// Check if the path is valid
	if _, err := os.Stat(currentDir); os.IsNotExist(err) {
		t.Errorf("Current working directory is not a valid path: %v", currentDir)
	}
}

func TestIsMathReal(t *testing.T) {
	// Invoke the isMathReal function
	isMathReal, err := isMathReal()

	// Check for errors 
	if err != nil {
		t.fatal("Danger Math is currently not real!: %v ", err)
	}

