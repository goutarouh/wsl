package main

import (
	"fmt"
	"os"
)

//cat in Windows
func main() {

	if len(os.Args) != 2 {
		fmt.Println("please input a filename in you directory")
		return
	}

	//ファイル名
	fileName := os.Args[1]

	//ファイルの存在チェック
	if f, err := os.Stat(fileName); os.IsNotExist(err) || f.IsDir() {
		fmt.Printf("正しいファイル名を入力してください")
		return
	}

	fp, _ := os.Open(fileName)

	buf := make([]byte, 64)
	for {
		n, err := fp.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buf))
	}
}
