// pointer.go ポインタについて
package main

import "fmt"

// ポインタを用いて入力された値を階乗に書き換える関数
func factorial(x *int) {
	for i := *x - 1; i > 0; i = i - 1 {
		*x = *x * i
	}
}

// ポインタを用いて入力されたVector3D型の値を書き換える関数
func changeVector3d(v *Vector3D) {
	// (*v).x としなくても、自動的に実体に代入してくれる
	v.x *= 2
	v.y *= 2
	v.z *= 2
}

// Vector3D 構造体　フィールド名が小文字で始まる...宣言したパッケージ内からのみ使用可能
//							 大文字で始まる...宣言したパッケージ以外からも使用可能
type Vector3D struct {
	x int
	y int
	z int
}

func main() {
	x := 5
	factorial(&x)
	fmt.Println(x)

	v := Vector3D{x: 1, y: 2, z: 3}
	fmt.Println(v)

	changeVector3d(&v)
	fmt.Println(v)
}
