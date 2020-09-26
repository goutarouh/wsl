package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

//ls in Windows
func main() {

	//カレントディレクトリ内のディレクトリ,ファイルのリストを取得
	files, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	//ファイル一覧を表示
	for _, file := range files {
		fmt.Println(file.Name())
	}

}