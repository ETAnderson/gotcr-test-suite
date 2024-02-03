package main

import (
	"os"
	"testing"
)

func getCurrentDirectory() (string, error) {
	return os.Getwd()
