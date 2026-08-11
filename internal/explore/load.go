// Package explore loads and analyzes Go source code to build the code
// reference graph — functions, types, and the relationships between them.
package explore

import (
	"fmt"

	"golang.org/x/tools/go/packages"
)

// Load loads and type-checks every package in the Go module rooted at dir.
// It uses go/packages, which wraps the real Go compiler's parser and type
// checker, so the result is compiler-accurate rather than guessed.

// A syntax error in one file does not fail the whole load — it's recorded
// on that package's Errors field instead. Use HasErrors to check.
func Load(dir string) ([]*packages.Package, error) {
	cfg := &packages.Config{
		Dir: dir,
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedDeps |
			packages.NeedTypes |
			packages.NeedTypesInfo |
			packages.NeedSyntax,
	}

	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return nil, fmt.Errorf("loading packages in %s: %w", dir, err)
	}

	return pkgs, nil
}

// HasErrors reports whether any loaded package failed to parse or
// type-check, printing the errors to stderr. Load itself still succeeds in
// this case — callers decide whether to skip the affected package.
func HasErrors(pkgs []*packages.Package) bool {
	return packages.PrintErrors(pkgs) > 0
}