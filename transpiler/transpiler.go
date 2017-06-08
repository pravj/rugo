// Package transpiler
package transpiler

import (
  "fmt"

  goast "go/ast"

  "github.com/pravj/rugo/ast"
  "github.com/pravj/rugo/program"
)

// transpile from IR to Go AST
func Transpile(node ast.Node, p *program.Program) {
  if node.Name == ":begin" {
    fmt.Println("iterate")
    for i := 0; i < len(node.Nodes); i++ {
      TranspileNode(true, node.Nodes[i], p)
    }
  } else {
    TranspileNode(true, p.Node, p)
  }
}

func TranspileNode(isRoot bool, node ast.Node, p *program.Program) goast.Node {
	switch node.Name {
  case ":true":
    return boolIdentifier(node)
  case ":false":
    return boolIdentifier(node)
  case ":int":
    return integerLiteral(node)
  case ":float":
    return floatLiteral(node)
  case ":str":
    return stringLiteral(node)
  case ":send":
    if ((len(node.Nodes) == 3) && (node.Nodes[0].Name == "nil") && (node.Nodes[1].Name == ":puts")) {
      stmt := printStmt(node.Nodes[2], p)
      if isRoot { p.AppendMainStatement(stmt) }

      return stmt
    }
  case ":lvasgn":
    stmt := localVarAssignStmt(node, p)
    if isRoot { p.AppendMainStatement(stmt) }

    return stmt
  case ":lvar":
    return localVariable(node)
  case ":if":
    stmt := ifStmt(node, p)
    if isRoot { p.AppendMainStatement(stmt) }

    return stmt
  default:
    panic(fmt.Sprintf("Unexpected node type %v", node.Name))
	}

  return nil
}
