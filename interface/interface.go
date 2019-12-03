package main

import "fmt"

// WritingUtensils インターフェース:メソッドだけ定義する
type WritingUtensils interface {
	PrintColor() string
}

// Pen ペンクラス
type Pen struct {
	Color string
}

// Paints 絵具クラス
type Paints struct {
	Color string
}

// PrintColor ...
func (p Pen) PrintColor() string {
	fmt.Println(p.Color)
	return p.Color
}

// PrintColor ...
func (p Paints) PrintColor() string {
	fmt.Println(p.Color)
	return p.Color
}

// IsRed 色が赤であればその旨を表示する
func IsRed(w WritingUtensils) {
	if w.PrintColor() == "red" {
		fmt.Println("この筆記用具は赤色を描きます.")
	} else {
		fmt.Println("この筆記用具は赤色を描きません.")
	}
}

func main() {
	orangePen := Pen{"orange"}
	redPaints := Paints{"red"}

	IsRed(orangePen)
	IsRed(redPaints)
}
