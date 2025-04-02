package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	src := []byte("a * (b + c)")

	// parser.ParseExprFromを使って解析する
	expr, err := parser.ParseExprFrom(fset, "", src, 0)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}

	// トークン位置を取得
	fmt.Println("ASTノード:", expr)
	fmt.Println("位置:", fset.Position(expr.Pos()))
}