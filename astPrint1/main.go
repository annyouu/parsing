package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	src := `
		package main
		import "fmt"
		var msg string
		func main() {
			msg = "Hello"
			name := "Gopher"
			fmt.Println(msg, name)
		}
	`

	// ファイルセット（位置情報管理）を作成
	fset := token.NewFileSet()

	// コードを解析し、ASTを生成
	f, err := parser.ParseFile(fset, "", src, parser.AllErrors)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	// ASTのダンプ(表示)
	ast.Print(fset, f)
}