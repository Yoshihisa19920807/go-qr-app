package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	msg := fmt.Sprintf("Hello %s\n", arg)

	// _とerrという２つの返り値が設定されている
	f, err := os.Create("hello.txt")
	if err != nil {
		fmt.Printf("failed to create file \n: %v", err)
		return
	}
	// 呼び出し元(ここではmain)がreturnをする際に実行される
	defer f.Close()

	_, err = f.WriteString(msg)
	if err != nil {
		fmt.Printf("failed to write message to file \n: %v", err)
		return
	}

	fmt.Println("Done!")
}
