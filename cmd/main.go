package main

import "path/filepath"

import (
	"fmt"
	"os"
	"github.com/gen2brain/beeep"
	gitutils "github.com/brijnandan/gitstreak/internal/git"
	"time"
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
	case "watch":
	    handleWatch()
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

    count, _ := gitutils.GetTodaysCommitCount()
    branch, _ := gitutils.GetCurrentBranch()
    lastCommit, _ := gitutils.GetLastCommitMessage()

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
	count, err := gitutils.GetTodaysCommitCount()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if count == "0" {
		err := sendReminder()

		if err != nil {
			fmt.Println("Failed to send notification:", err)
			return
		}

		fmt.Println("Reminder sent")
		return
	}

	fmt.Printf("Streak safe (%s commits today)\n", count)
}

func handleWatch() {
	fmt.Println("Forge watcher started")

	notified := false

	for {
		count, err := gitutils.GetTodaysCommitCount()

		if err == nil {

			if count == "0" && !notified {
				sendReminder()
				notified = true
			}

			if count != "0" {
				notified = false
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func printHelp() {
	fmt.Println("Forge")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  forge status")
	fmt.Println("  forge remind")
	fmt.Println("  forge watch")
	fmt.Println("  forge help")
}
func getRepoName() string {
	return filepath.Base(getCurrentDir())
}

func getCurrentDir() string {
	dir, _ := os.Getwd()
	return dir
}

func sendReminder() error {
	return beeep.Notify(
		"Forge",
		"No commits today. Your streak is at risk.",
		"",
	)
}
