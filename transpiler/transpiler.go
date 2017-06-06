// Package transpiler
package transpiler

import (
  "fmt"

  goast "go/ast"
  gotoken "go/token"

  "github.com/pravj/rugo/ast"
  "github.com/pravj/rugo/program"
)

// transpile from IR to Go AST
func Transpile(node ast.Node, p *program.Program) {
	TranspileNode(p.Node, p)
}

func integerLiteral(node ast.Node) *goast.BasicLit {
  return &goast.BasicLit{
    Kind: gotoken.INT,
    Value: node.Nodes[0].Name,
  }
}

func stringLiteral(node ast.Node) *goast.BasicLit {
  return &goast.BasicLit{
    Kind: gotoken.STRING,
    Value: node.Nodes[0].Name,
  }
}

func floatLiteral(node ast.Node) *goast.BasicLit {
  return &goast.BasicLit{
    Kind: gotoken.FLOAT,
    Value: node.Nodes[0].Name,
  }
}

func boolIdentifier(node ast.Node) *goast.Ident {
  return goast.NewIdent(node.Name[1:])
}

func localVarAssignStmt(node ast.Node, p *program.Program) goast.Stmt {
  if len(node.Nodes) == 2 {
    tokenType := gotoken.DEFINE

    stmt := &goast.AssignStmt{
      Lhs: []goast.Expr{
        &goast.Ident{
          Name: node.Nodes[0].Name[1:],
        },
      },
      Tok: tokenType,
      Rhs: []goast.Expr{TranspileNode(node.Nodes[1], p).(goast.Expr)},
    }

    return stmt
  } else {
    panic("Incorrect local-variable-assign statement")
  }
}

func TranspileNode(node ast.Node, p *program.Program) goast.Node {
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
  case ":lvasgn":
    p.AppendMainStatement(localVarAssignStmt(node, p))
  default:
    panic(fmt.Sprintf("Unexpected node %v", node.Name))
	}

  return nil
}
