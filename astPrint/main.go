package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// token.FileSetの作成(位置情報管理)
	fset := token.NewFileSet()

	// 式 v + 1を解析してASTを生成
	expr, err := parser.ParseExpr("v + 1")
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}

	// ASTの全構造をダンプ(表示)する
	ast.Print(fset, expr)
}