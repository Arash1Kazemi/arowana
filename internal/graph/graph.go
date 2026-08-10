// Package graph defines the shared types both the gitlog and explore

package graph

// Node is one point in a graph a commit (gitlog), or a symbol from go/types or a language server (explore)
type Node struct {
	ID       string         `json:"id"`
	Label    string         `json:"label"`
	Kind     string         `json:"kind"` // "commit" | "function" | "method" | "type" | "interface" | "var" | "const" | 
	Location *Location      `json:"location,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

// Location points at a spot in a source file
type Location struct {
	File   string `json:"file"`
	Line   int    `json:"line"`
	Column int    `json:"column"`
}

type Edge struct {
	From     string         `json:"from"`
	To       string         `json:"to"`
	Kind     string         `json:"kind"` // "parent" | "merge" | "calls" | "implements" | "embeds" | "uses-type"
	Metadata map[string]any `json:"metadata,omitempty"`
}

// Graph is what both engines hand to the frontend.
type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}