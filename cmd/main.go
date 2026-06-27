package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/beeep"

	configpkg "github.com/brijnandan/gitstreak/internal/config"

	gitutils "github.com/brijnandan/gitstreak/internal/git"

	"strconv"
)

const Version = "0.1.0"

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
    case "version":
	  handleVersion()
    case "add":
	  handleAdd()
    case "help":
	  printHelp()
	case "remove":
	  handleRemove()
	case "list":
	  handleList()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printHelp()
	}
}

func handleStatus() {
	cfg, err := configpkg.LoadConfig()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	totalCommits := 0

	for _, repo := range cfg.Repositories {
		countStr, err := gitutils.GetTodaysCommitCountForRepo(repo)

		if err != nil {
			continue
		}

		count, _ := strconv.Atoi(countStr)
		totalCommits += count
	}

	status := "AT RISK"

	if totalCommits > 0 {
		status = "SAFE"
	}

	fmt.Println("Forge")
	fmt.Println()

	fmt.Printf("%-22s %d\n", "Tracked Repositories", len(cfg.Repositories))
	fmt.Printf("%-22s %d\n", "Commits Today", totalCommits)
	fmt.Printf("%-22s %s\n", "Status", status)
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
	fmt.Println("  forge version")
	fmt.Println("  forge list")
}

func sendReminder() error {
	return beeep.Notify(
		"Forge",
		"No commits today. Your streak is at risk.",
		"",
	)
}

func handleVersion() {
	fmt.Printf("Forge %s\n", Version)
}

func handleAdd() {
	err := configpkg.EnsureConfigDir()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	path, err := os.Getwd()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = configpkg.AddRepository(path)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Repository added:")
	fmt.Println(path)
}

func handleList() {
	cfg, err := configpkg.LoadConfig()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(cfg.Repositories) == 0 {
		fmt.Println("No repositories configured")
		return
	}

	fmt.Println("Tracked Repositories")
	fmt.Println()

	for i, repo := range cfg.Repositories {
		fmt.Printf("%d. %s\n", i+1, repo)
	}
}

