package explore

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "go.mod", "module fixture\n\ngo 1.26\n")
	writeFile(t, dir, "main.go", `package main

func add(a, b int) int { return a + b }

func main() { add(1, 2) }`)

	pkgs, err := Load(dir)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if len(pkgs) != 1 {
		t.Fatalf("got %d packages, want 1", len(pkgs))
	}
	if HasErrors(pkgs) {
		t.Fatal("fixture package has unexpected errors")
	}
}

func writeFile(t *testing.T, dir, name, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
		t.Fatalf("writing %s: %v", name, err)
	}
}