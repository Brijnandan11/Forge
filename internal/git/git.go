package git

import (
	"os/exec"
	"strings"
)

func GetTodaysCommitCount() (string, error) {
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

func GetCurrentBranch() (string, error) {
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

func GetLastCommitMessage() (string, error) {
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

func GetTodaysCommitCountForRepo(repoPath string) (string, error) {
	cmd := exec.Command(
		"git",
		"rev-list",
		"--count",
		"--since=midnight",
		"HEAD",
	)

	cmd.Dir = repoPath

	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}
