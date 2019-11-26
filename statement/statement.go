// statement.go ステートメントについて

package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// memo:1 defer...関数のスコープを抜ける時に実行
	defer fmt.Println("おしまい")
	// memo:2 deferは後ろに書いた式から実行される。
	defer fmt.Println("今日は ")

	// memo:3 switch文で変数定義を同時に行う スコープはこのswitch文のみ
	switch nation := "Japan"; nation {
	case "Usa":
		fmt.Println("USA")
	case "Gbr":
		fmt.Println("GBR")
	default:
		fmt.Println(nation)
	}

	// memo:4 switch文に直接式を入れる
	switch {
	case time.Now().Hour() < 12:
		fmt.Println("AM")
	default:
		fmt.Println("PM")
	}

	// memo:5 ファイル読み込み, エラーハンドリング, ログ
	file, err := os.Open("./sample.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// memo:6 deferを用いたファイルのクローズ
	defer file.Close()

	// memo:7 ファイルの内容を表示
	fileData := make([]byte, 100)
	file.Read(fileData)
	fmt.Println(string(fileData))
}
