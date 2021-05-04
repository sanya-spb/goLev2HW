package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fmt.Println(fileFuncAnaliser("./srs_02/test1.go", "foo"))
}

func fileFuncAnaliser(fileName, funcName string) (int32, error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, fileName, nil, 0)
	if err != nil {
		return 0, nil
	}
	var count int32
	for _, decl := range astFile.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		count += findInStatement(funcDecl.Body, funcName)

		if funcDecl.Name.String() == funcName {
			count++
		}
	}
	return count, nil
}

func findInStatement(stmt ast.Stmt, funcName string) (count int32) {
	switch v := stmt.(type) {
	case *ast.BlockStmt:
		for _, block := range v.List {
			count += findInStatement(block, funcName)
		}
	case *ast.ExprStmt:
		count += findInExpr(v.X, funcName)
	case *ast.IfStmt:
		count += findInExpr(v.Cond, funcName)
		count += findInStatement(v.Body, funcName)
		if stmt.(*ast.IfStmt).Else != nil {
			count += findInStatement(v.Else, funcName)
		}
	case *ast.ForStmt:
		count += findInStatement(v.Init, funcName)
		count += findInStatement(v.Body, funcName)
		count += findInExpr(v.Cond, funcName)
	case *ast.AssignStmt:
		for _, assign := range v.Rhs {
			count += findInExpr(assign, funcName)
		}
	case *ast.RangeStmt:
		count += findInExpr(v.X, funcName)
		count += findInStatement(v.Body, funcName)
	case *ast.ReturnStmt:
		for _, ret := range v.Results {
			count += findInExpr(ret, funcName)
		}
	}
	return count
}

func findInExpr(e ast.Expr, funcName string) (count int32) {
	switch v := e.(type) {
	case *ast.BinaryExpr:
		count += findInExpr(v.X, funcName)
		count += findInExpr(v.Y, funcName)
	case *ast.Ident:
		if v.Name == funcName {
			count++
		}
	case *ast.CallExpr:
		count += findInExpr(v.Fun, funcName)
		for _, arg := range v.Args {
			count += findInExpr(arg, funcName)
		}
	case *ast.SelectorExpr:
		count += findInExpr(v.Sel, funcName)
	case *ast.BasicLit:
		if v.Value == funcName {
			count++
		}
	}
	return count
}
