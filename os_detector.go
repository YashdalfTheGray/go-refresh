package main

import (
	"fmt"
	"runtime"
)

// DetectOS prints out what operating system this binary is running on
func DetectOS() {
	var message string
	switch os := runtime.GOOS; os {
	case "darwin":
		message = "Running on macOS"
		break
	case "windows":
		message = "Running on Windows"
		break
	case "linux":
		message = "Running on linux"
		break
	default:
		message = "Running on a different OS"
		break
	}

	fmt.Println(message)
}
