package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	repoPath := "."
	if len(os.Args) > 2 {
		repoPath = os.Args[2]
	}

	switch os.Args[1] {
	case "gitlog":
		runGitlog(repoPath)
	case "explore":
		runExplore(repoPath)
	case "both":
		runBoth(repoPath)
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: arowana <gitlog|explore|both> [repo-path]")
}

// Stubs for now — Milestone 1 fills in runExplore first.
func runGitlog(repoPath string)  { fmt.Println("gitlog: not implemented yet —", repoPath) }
func runExplore(repoPath string) { fmt.Println("explore: not implemented yet —", repoPath) }
func runBoth(repoPath string)    { fmt.Println("both: not implemented yet —", repoPath) }
