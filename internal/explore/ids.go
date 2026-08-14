package explore

import (
	"go/token"
	"go/types"
	"strings"

	"github.com/Arash1Kazemi/arowana/internal/graph"
	"golang.org/x/tools/go/packages"
)

// symbolID returns a stable, unique ID for a declared symbol.
//
//	sym:<pkgpath>.<name>          functions, types, vars, consts
//	sym:<pkgpath>.<recv>.<name>   methods
func symbolID(obj types.Object) string {
	if obj == nil || obj.Pkg() == nil {
		return ""
	}

	var b strings.Builder
	b.WriteString("sym:")
	b.WriteString(obj.Pkg().Path()) // Use the import path, not the package name, to avoid collisions between packages with the same name.
	b.WriteString(".")
	b.WriteString(".")
	// receiverName helps disambiguate methods with the same name on different types.
	//  Without the receiver, Circle.Area and Square.Area both become
	// "sym:interfaces.Area" and collapse into one node.
	if recv := receiverName(obj); recv != "" {
		b.WriteString(recv)
	}
	b.WriteString(obj.Name())
	return b.String()
}

// receiverName returns the name of a method's receiver type, or "" if obj
// is not a method. Pointer and value receivers are treated the same.
// is this symbol a method, and on what type?
func receiverName(obj types.Object) string {
	fn, ok := obj.(*types.Func)
	if !ok {
		return ""
	}
	sig, ok := fn.Type().(*types.Signature)
	if !ok || sig.Recv() == nil {
		return ""
	}
	return typeName(sig.Recv().Type())
}

// typeName unwraps pointers to get the underlying named type's name
// Wrapper over namedType that returns the name as a string instead of the object.
func typeName(t types.Type) string {
	if named := namedType(t); named != nil {
		return named.Obj().Name()
	}
	return ""
}

// namedType unwraps a pointer (if any) and returns the named type behind it
// or nil if there isn't one (slices, maps, unnamed structs, ...).
func namedType(t types.Type) *types.Named {
	// Without this every pointer-receiver method
	// gets an empty receiver name and silently collides.
	if ptr, ok := t.(*types.Pointer); ok {
		t = ptr.Elem()
	}
	named, _ := t.(*types.Named)
	return named
}

// position converts a token.Pos into a graph.Location, or nil when the
// position is unknown (compiler-synthesized symbols have no source).
func position(pkg *packages.Package, pos token.Pos) *graph.Location {
	if pkg.Fset == nil || !pos.IsValid() {
		return nil
	}
	p := pkg.Fset.Position(pos)
	return &graph.Location{File: p.Filename, Line: p.Line, Column: p.Column}
}

// Builds the set of package paths that actually loaded.
// localPackages returns the set of package paths that were explicitly
// loaded. Symbols outside this set are skipped — the graph describes this
// project, not everything it imports.
func localPackages(pkgs []*packages.Package) map[string]bool {
	local := make(map[string]bool, len(pkgs))
	for _, pkg := range pkgs {
		local[pkg.PkgPath] = true // to filter fmt.Println add stdlib nodes and edges on eatch call
	}
	return local
}

// isLocal reports whether obj was declared in one of the loaded packages.
func isLocal(obj types.Object, local map[string]bool) bool {
	return obj != nil && obj.Pkg() != nil && local[obj.Pkg().Path()]
}
