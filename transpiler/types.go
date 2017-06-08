// Package transpiler
package transpiler

import (
  goast "go/ast"
  gotoken "go/token"

  "github.com/pravj/rugo/ast"
)

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

func localVariable(node ast.Node) *goast.Ident {
  return goast.NewIdent(node.Nodes[0].Name[1:])
}
