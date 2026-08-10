package main

import (
	"fmt"
	"log"
	"net/http"
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

// Temporary — just enough to prove the Vite proxy reaches the Go server.
// Real endpoints land later .
func runExplore(repoPath string) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok","repo":"` + repoPath + `"}`))
	})
	log.Println("explore server listening on :8080 —", repoPath)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// Stubs for now.
func runGitlog(repoPath string) { fmt.Println("gitlog: not implemented yet —", repoPath) }
func runBoth(repoPath string)   { fmt.Println("both: not implemented yet —", repoPath) }