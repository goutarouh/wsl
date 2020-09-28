package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

//ls in Windows
func main() {

	//flagの引数の設定
	var (
		l = flag.Bool("l", false, "display all information of the files such as file size")
	)

	flag.Parse()

	//カレントディレクトリ内のディレクトリ,ファイルのリストを取得
	files, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	//ファイル一覧を表示
	for _, file := range files {
		if *l {
			fmt.Println(file.Size(), file.Name())
		} else {
			fmt.Print(file.Name(), " ")
		}
	}

}
