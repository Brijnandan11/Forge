package main

import "path/filepath"

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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
	_, err := os.Stat(".git")

	if err != nil {
		fmt.Println("✗ Not a Git repository")
		return
	}

	count, err := getTodaysCommitCount()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("✓ Git repository detected")
	fmt.Printf("Today's Commits: %d\n", count)

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

func getTodaysCommitCount() (int, error) {
	cmd := exec.Command(
		"git",
		"log",
		"--since=today",
		"--oneline",
	)

	output, err := cmd.Output()

	if err != nil {
		return 0, err
	}

	lines := strings.Split(
		strings.TrimSpace(string(output)),
		"\n",
	)

	if len(lines) == 1 && lines[0] == "" {
		return 0, nil
	}

	return len(lines), nil
}

func getRepoName() string {
	return filepath.Base(getCurrentDir())
}

func getCurrentDir() string {
	dir, _ := os.Getwd()
	return dir
}