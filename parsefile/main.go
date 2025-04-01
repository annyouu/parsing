package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"go/ast"
)

func main() {
	// 解析するソースコード
	const src = `
		package main
		var v = 100
		func main() {
			println(v+1)
		}`
	
	// ファイルセットを作成
	fset := token.NewFileSet()

	// parser.ParseFileを使って構文解析
	f, err := parser.ParseFile(fset, "my.go", src, 0)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}

	// パッケージ名の取得
	fmt.Println("パッケージ名:", f.Name.Name)

	// 定義された識別子を表示
	for _, decl := range f.Decls { // 全ての宣言を取得
		if genDecl, ok := decl.(*ast.GenDecl); ok { // 一般的な宣言(変数、定数、型、import)か確認
			for _, spec := range genDecl.Specs { // 各宣言(変数、定数など)を取得する
				if valueSpec, ok := spec.(*ast.ValueSpec); ok { // 変数or定数か確認する
					for _, name := range valueSpec.Names { // 変数名リストを取得
						fmt.Println("変数:", name.Name) // 変数名を出力
					}
				}
			}
		}
	}
}