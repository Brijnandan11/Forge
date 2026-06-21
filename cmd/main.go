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
	wd, _ := os.Getwd()
    repo := filepath.Base(wd)

    count, _ := getTodaysCommitCount()
    branch, _ := getCurrentBranch()
    lastCommit, _ := getLastCommitMessage()

    fmt.Println("GitStreak")
    fmt.Println()

    fmt.Printf("%-15s %s\n", "Repository", repo)
    fmt.Printf("%-15s %s\n", "Branch", branch)
    fmt.Printf("%-15s %s\n", "Commits Today", count)
    fmt.Printf("%-15s %s\n", "Last Commit", lastCommit)

    if count == "0" {
	   fmt.Printf("%-15s %s\n", "Status", "AT RISK")
    } else {
	    fmt.Printf("%-15s %s\n", "Status", "SAFE")
   }

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

func getTodaysCommitCount() (string, error) {
	cmd := exec.Command(
		"git",
		"rev-list",
		"--count",
		"--since=midnight",
		"HEAD",
	)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func getRepoName() string {
	return filepath.Base(getCurrentDir())
}

func getCurrentDir() string {
	dir, _ := os.Getwd()
	return dir
}

func getLastCommitMessage() (string, error) {
	cmd := exec.Command(
		"git",
		"log",
		"-1",
		"--pretty=%s",
	)

	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func getCurrentBranch() (string, error) {
	cmd := exec.Command(
		"git",
		"branch",
		"--show-current",
	)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}