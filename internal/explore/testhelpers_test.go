package explore

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/packages"
)

// loadFixture loads one of the testdata modules by directory name and
// fails the test immediately if it doesn't load or type-check cleanly.
// Every node/edge extraction test starts by calling this, so a broken
// fixture is caught here instead of producing confusing failures later.
func loadFixture(t *testing.T, name string) []*packages.Package {
	t.Helper()

	pkgs, err := Load(filepath.Join("testdata", name))
	if err != nil {
		t.Fatalf("loading fixture %s: %v", name, err)
	}
	if HasErrors(pkgs) {
		t.Fatalf("fixture %s failed to type-check", name)
	}
	return pkgs
}

// idSet turns a slice into a set of string keys, using the given id
// function to compute each key. Used by tests that only need to check
// "did we extract this specific node/edge", not compare full structs.
func idSet[T any](items []T, id func(T) string) map[string]bool {
	set := make(map[string]bool, len(items))
	for _, item := range items {
		set[id(item)] = true
	}
	return set
}
