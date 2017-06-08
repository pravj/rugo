// Package transpiler
package transpiler

import (
  "fmt"

  goast "go/ast"
  gotoken "go/token"

  "github.com/pravj/rugo/ast"
  "github.com/pravj/rugo/program"
)

func localVarAssignStmt(node ast.Node, p *program.Program) goast.Node {
  if len(node.Nodes) == 2 {
    tokenType := gotoken.DEFINE

    stmt := &goast.AssignStmt{
      Lhs: []goast.Expr{
        &goast.Ident{
          Name: node.Nodes[0].Name[1:],
        },
      },
      Tok: tokenType,
      Rhs: []goast.Expr{TranspileNode(false, node.Nodes[1], p).(goast.Expr)},
    }

    return stmt
  } else {
    panic("Incorrect local-variable-assign statement")
  }
}

func printStmt(node ast.Node, p *program.Program) goast.Node {
  stmt := &goast.ExprStmt{
    X: &goast.CallExpr{
      Fun: &goast.SelectorExpr{
        X: goast.NewIdent("fmt"),
        Sel: goast.NewIdent("Println"),
      },
      Args: []goast.Expr{TranspileNode(false, node, p).(goast.Expr)},
    },
  }

  // import the "fmt" package
  p.AddImport("fmt")

  return stmt
}

func ifStmt(node ast.Node, p *program.Program) goast.Stmt {
  stmt := &goast.IfStmt{
    Cond: TranspileNode(false, node.Nodes[0], p).(goast.Expr),
    Body: &goast.BlockStmt{
      List: ifStmtBody(node.Nodes[1], p),
    },
  }

  return stmt
}

func ifStmtBody(node ast.Node, p *program.Program) []goast.Stmt {
  var stmts []goast.Stmt

  if node.Name == ":begin" {
    fmt.Println("begin in if")
    for i := 0; i < len(node.Nodes); i++ {
      stmts = append(stmts, TranspileNode(false, node.Nodes[i], p).(goast.Stmt))
    }
  } else {
    stmts = append(stmts, TranspileNode(false, node, p).(goast.Stmt))
  }

  return stmts
}
