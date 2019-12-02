package main

import "math"

import "fmt"

// Vector3D 構造体　フィールド名が小文字で始まる...宣言したパッケージ内からのみ使用可能
//							 大文字で始まる...宣言したパッケージ以外からも使用可能
type Vector3D struct {
	x int
	y int
	z int
}

// SizeofVector3D メソッド...Vector3D.SizeofVector3D()の形で呼び出せるようになる.
func (v Vector3D) SizeofVector3D() float64 {
	return math.Sqrt(math.Pow(float64(v.x), 2) + math.Pow(float64(v.y), 2) + math.Pow(float64(v.z), 2))
}

// ChangeXZAxis x座標とz座標の値を入れ替える(ポインタ レシーバー)
func (v *Vector3D) ChangeXZAxis() {
	v.x, v.y = v.z, v.x
}

// New コンストラクタ
func New(x int, y int, z int) *Vector3D {
	return &Vector3D{x, y, z}
}

func main() {
	vec := New(3, 4, 5)

	fmt.Println(*vec)

	fmt.Println(vec.SizeofVector3D())

	vec.ChangeXZAxis()
	fmt.Println(*vec)
}
