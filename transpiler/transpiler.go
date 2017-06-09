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
    return sendStmt(isRoot, node, p)
  case ":lvasgn":
    return localVarAssignStmt(isRoot, node, p)
  case ":lvar":
    return localVariable(node)
  case ":if":
    return ifStmt(isRoot, node, p)
  case ":def":
    return def
  default:
    panic(fmt.Sprintf("Unexpected node type %v", node.Name))
	}

  return nil
}
