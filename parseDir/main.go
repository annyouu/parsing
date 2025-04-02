package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getGOROOT() (string, error) {
	cmd := exec.Command("go", "env", "GOROOT")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func main() {
	// GOROOTの取得
	goroot, err := getGOROOT()
	if err != nil {
		fmt.Println("GOROOTの取得に失敗:", err)
		return
	}

	// token.FileSetを作成(ファイル位置情報の管理)
	fset := token.NewFileSet()

	// fmtパッケージのディレクトリパスを取得
	path := filepath.Join(goroot, "src", "fmt")

	// format.goだけを解析対象にするフィルタ
	filter := func(info os.FileInfo) bool {
		return info.Name() == "format.go"
	}

	// fmtパッケージをディレクトリ単位で解析
	pkgs, err := parser.ParseDir(fset, path, filter, 0)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}

	// 解析結果を出力
	for name, pkg := range pkgs {
		fmt.Printf("======= パッケージ: %s =======\n", name)
		for fname := range pkg.Files {
			fmt.Println("解析したファイル:", fname)
		}
	}
}