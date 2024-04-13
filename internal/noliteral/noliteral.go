package noliteral

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "noliteral",
	Doc: "noliteral finds calls to x.Exec with string literal arg",
	Run: run,
	Requires: []*analysis.Analyzer{ inspect.Analyzer },
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{(*ast.CallExpr)(nil)}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.CallExpr:
			if isGetConst(n) {
				if isStringLiteral(n.Args[0]) {
					l := n.Args[0].(*ast.BasicLit).Value
					l = l[1:len(l)-1]
					pass.Reportf(n.Args[0].Pos(), "Use x.%s instead of string literal", l)
				}
			}
		}
	})
	return nil, nil
}

func isGetConst(n *ast.CallExpr) bool {
	f, ok := n.Fun.(*ast.SelectorExpr)
	if !ok || f == nil {
		return false
	}

	pkg, ok := f.X.(*ast.Ident)
	if !ok || pkg == nil || pkg.Name != "x" {
		return false
	}

	fName := f.Sel
	if fName == nil || fName.Name != "Exec" {
		return false
	}

	return true
}

func isStringLiteral(n ast.Expr) bool {
	if _, ok := n.(*ast.BasicLit); !ok {
		return false
	}
	return true
}