// define.go 定義について

package main

import (
	"fmt"
	"strconv"
)

// init main()の実行前に呼び出される。
func init() {
	fmt.Println("ようこそ。")
}

// Greet greet: 各国の挨拶 ex. Hello.
//		 name : 名前
func Greet(greet string) func(name string) string {
	return func(name string) string {
		return greet + " " + name
	}
}

// Multiplication 全ての整数値を乗算する
func Multiplication(intValue ...int) (amount int) {
	amount = 1
	for _, value := range intValue {
		amount *= value
	}
	return
}

func main() {
	fmt.Println("メイン関数だよ。")

	floatValue := 1.34
	// memo:1 %Tで型を表示する
	fmt.Printf("floatValue Type is: %T\n", floatValue)
	// memo:2 float32で変数を定義する
	var float32Value float32 = 1.34
	fmt.Printf("float32Value Type is: %T\n", float32Value)

	// memo:3 文字列型の数値を数値として扱う
	numberString := "12345"

	// memo:4 '_'...戻り値として用意しなければならないものの、使用しない値を_で置き換える
	// ex. strconv.Atoi(string)は戻り値としてintとerrorを持つが、errorは使わないので_で置きかえる
	number, _ := strconv.Atoi(numberString)
	fmt.Printf("number:%T - %d\n", number, number)

	// memo: 5 クロージャ
	japanese := Greet("こんにちは.")
	fmt.Println(japanese("Taro Yamada"))
	english := Greet("Hello.")
	fmt.Println(english("Michael Johnson"))

	// memo: 6 可変長引数
	fmt.Println("Multiplication Result:", Multiplication(1, 2, 3, 4, 5))
	// memo: 7 配列渡し
	valueArray := []int{10, 20, 30, 40, 50}
	fmt.Println("Multiplication Result:", Multiplication(valueArray...))
}
