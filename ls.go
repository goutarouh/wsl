package main

import (
	"flag"
	"fmt"
	"github.com/gookit/color"
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

	//テーブルヘッダの表示
	if *l {
		fmt.Printf("%-4s %s\n", "サイズ", "ファイル名")
	}

	//ファイル一覧を表示
	for _, file := range files {
		if file.Name()[0] == '.' {
			continue
		}
		if *l {
			if file.IsDir() {
				color.Cyan.Printf("%s\n", file.Name())
			} else {
				fmt.Printf("%-4d %s\n", file.Size(), file.Name())
			}
		} else {
			fmt.Print(file.Name(), " ")
		}
	}

}
