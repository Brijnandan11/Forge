package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "status":
		handleStatus()
	case "remind":
		handleRemind()
	case "help":
		printHelp()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printHelp()
	}
}

func handleStatus() {
	fmt.Println("Checking repository...")
}

func handleRemind() {
	fmt.Println("Checking today's commits...")
}

func printHelp() {
	fmt.Println("GitStreak")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  gitstreak status")
	fmt.Println("  gitstreak remind")
	fmt.Println("  gitstreak help")
}