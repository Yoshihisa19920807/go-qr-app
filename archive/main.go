// このファイルのパッケージ名を宣言
package main

// パッケージを読み込み
import (
	"flag"
	"fmt"
)

func main() {
	// fmt.Println("Hello world 🍣")
	flag.Parse()
	// 「:=」で値の宣言と代入
	// 「()」の数字で何番目の引数を使うかを指定
	arg := flag.Arg(0)
	fmt.Printf("Hello %s\n", arg)
}
