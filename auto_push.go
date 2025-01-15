package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	gitRepoPath   = "."
	fileToUpdate  = "activity_log.txt" // File to update
	commitMessage = "Automated update via Go program"
)

// updateTextFile appends the current timestamp to a specified file.
func updateTextFile() error {
	filePath := filepath.Join(gitRepoPath, fileToUpdate)

	// Open the file in append mode, creating it if it doesn't exist
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Write the current timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	_, err = file.WriteString(fmt.Sprintf("Update at %s\n", timestamp))
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	fmt.Printf("Updated file: %s\n", filePath)
	return nil
}

// runGitCommand executes a Git command in the repository's directory.
func runGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Dir = gitRepoPath

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git command failed: %s, output: %s", err, string(output))
	}

	fmt.Printf("Git command output: %s\n", string(output))
	return nil
}

// gitCommitAndPush stages, commits, and pushes changes to the remote repository.
func gitCommitAndPush() error {
	// Stage changes
	if err := runGitCommand("add", "."); err != nil {
		return fmt.Errorf("failed to stage changes: %w", err)
	}

	// Commit changes
	if err := runGitCommand("commit", "-m", commitMessage); err != nil {
		return fmt.Errorf("failed to commit changes: %w", err)
	}

	// Push changes
	if err := runGitCommand("push"); err != nil {
		return fmt.Errorf("failed to push changes: %w", err)
	}

	fmt.Println("Changes committed and pushed successfully.")
	return nil
}

func main() {
	// Step 1: Update the text file
	if err := updateTextFile(); err != nil {
		fmt.Printf("Error updating text file: %v\n", err)
		return
	}

	// Step 2: Commit and push changes
	if err := gitCommitAndPush(); err != nil {
		fmt.Printf("Error during git operations: %v\n", err)
	}
}
