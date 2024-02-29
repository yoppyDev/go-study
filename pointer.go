package main

import (
	"fmt"
)

func main () {
	var n int = 1

	// 値を表示
	fmt.Println(n)
	// ポインタの値(アドレスを表示)
	fmt.Println(&n)
	// ポインタが表す値を表示
	fmt.Println(*&n)

	// メモリの確保
	var p *int
	fmt.Println(&p)
	fmt.Println(p)

	// 構造体の初期化
	Vertex := struct {
		X, Y int
		ms   string
	}{
		X:  1,
		Y:  2,
		ms: "test",
	}

	fmt.Println(Vertex.X, Vertex.Y, Vertex.ms)
}