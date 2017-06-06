// Package ast
package ast

// Node represents a node in Intermediate Representation's AST
type Node struct {
	Name  string
	Nodes []Node
}

// NewNode returns a new Node instance (*Node)
func NewNode() Node {
	return Node{}
}
