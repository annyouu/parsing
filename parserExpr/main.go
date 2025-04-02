package main

import (
	"fmt"
	"go/parser"
	"go/ast"
)

func main() {
	// 式の解析
	expr, err := parser.ParseExpr("x + 42")
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}

	// 抽象構文木(AST)の確認
	fmt.Printf("解析結果: %#v\n", expr)

	// 型アサーションを使って演算子を取得
	if binaryExpr, ok := expr.(*ast.BinaryExpr); ok {
		fmt.Println("左辺:", binaryExpr.X)
		fmt.Println("演算子:", binaryExpr.Op)
		fmt.Println("右辺:", binaryExpr.Y)
	}
}