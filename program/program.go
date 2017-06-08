// Package program
package program

import (
	"bytes"
	"fmt"
	"strconv"
	"go/format"

	goast "go/ast"
	goparser "go/parser"
	gotoken "go/token"

	"github.com/pravj/rugo/ast"
	"github.com/pravj/rugo/utils"
)

// Program represents the metadata about the final Go program
type Program struct {
	// name of the generated package
	PackageName string

	// node containing the entire program code
	Node ast.Node

	// set of source files
	// single file in our case, will be multiple for a package
	FileSet *gotoken.FileSet

	// the source (single) Go file
	File *goast.File

	// statements having the global scope (for the main function)
	MainStatements []goast.Stmt

	// map from an identifier to its scope
	// helpful in case of Ruby's assignment to check if a name is already defined
	SymbolTable map[string]string

	// imported path for the program
	Imports []string
}

// NewProgram returns a new program instance (*Program)
func NewProgram(packageName string, node ast.Node) *Program {
	fileSet := gotoken.NewFileSet()

	return &Program{
		PackageName: packageName,
		Node:        node,
		FileSet:     fileSet,
	}
}

// String returns the formated string representation of the code
func (p *Program) String() string {
	packageSignature := fmt.Sprintf("package %v", p.PackageName)
	f, err := goparser.ParseFile(p.FileSet, "", packageSignature, 0)

	utils.CheckError(err)
	p.File = f

	// main function
	p.File.Decls = append(p.File.Decls, &goast.FuncDecl{
		Name: goast.NewIdent("main"),
		Type: &goast.FuncType{
			Results: nil,
		},
		Body: &goast.BlockStmt{
			List: p.MainStatements,
		},
	})

	// add grouped imports
	importDecl := &goast.GenDecl{
		Tok:    gotoken.IMPORT,
		Lparen: 1,
	}

	for _, quotedImportPath := range p.Imports {
		importSpec := &goast.ImportSpec{
			Path: &goast.BasicLit{
				Kind:  gotoken.IMPORT,
				Value: quotedImportPath,
			},
		}

		importDecl.Specs = append(importDecl.Specs, importSpec)
	}

	p.File.Decls = append([]goast.Decl{importDecl}, p.File.Decls...)

	// use the backbone of "gofmt" to get a string representation from AST
	var buf bytes.Buffer
	format.Node(&buf, p.FileSet, p.File)

	return buf.String()
}

// AppendMainStatement adds a new statement to the main function.
func (p *Program) AppendMainStatement(stmt goast.Node) {
	p.MainStatements = append(p.MainStatements, stmt.(goast.Stmt))
}

// AddImport adds an import path if it's not in the import list.
func (p *Program) AddImport(importPath string) {
	quotedImportPath := strconv.Quote(importPath)

	for _, i := range p.Imports {
		// ignore if the path is in the list
		if i == quotedImportPath {
			return
		}
	}

	p.Imports = append(p.Imports, quotedImportPath)
}
