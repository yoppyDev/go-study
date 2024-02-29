// 宣言 main は、main パッケージ内でのみ使用可能です
package main

// 使わないモジュールはビルド時にエラーになる
import (
	"fmt"
	"os/user"
	"strings"
	"time"
	"strconv"
)

// 変数は関数外で定義できる
var (
	i   int = 1
	f64     = 1.2
	str     = "Hello World"
)

const (
	// 定数
	i2 int = 2
)

// init関数はmain関数よりも先に実行される
func init() {
	fmt.Println("Init!")
}

func add(x, y int) (sum int) {
	sum = x + y
	return
}

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	fmt.Println("Hello world!", time.Now())
	fmt.Println(user.Current())
	fmt.Println(i)

	// 宣言した変数は関数内でのみ使用可能
	xi := 1
	fmt.Println(xi)
	// fmt.Printf("%T\n", xi)で型を表示
	fmt.Printf("%T\n", i2)
	fmt.Printf("type=%T value=%v\n", i2, i2)

	// 文字列置き換え
	fmt.Println(strings.Replace(str, "World", "Go", 1))

	// 型変換
	var s string = "14"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)

	// 配列はリサイズできないが、スライスはリサイズできる

	// 配列
	var arr [2]int
	fmt.Println(arr)

	// スライス
	var a []int
	a = []int{1, 2, 3}
	a = append(a, 4)
	fmt.Println(a)

	var b []int
	// make関数でスライスを作成
	b = make([]int, 3, 5)
	// キャパシティとは、スライスの要素数がいくつまで増やせるかを示す
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	// len=3 cap=5 value=[0 0 0]

	var c []int
	// キャパシティを指定しない
	// [0 0 0 0 0 0 1 2 3 4]
	c = make([]int, 5)
	// キャパシティを指定
	// [0 1 2 3 4]
	c = make([]int, 0, 5)

	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

	sum := add(1, 2)
	fmt.Println(sum)

	// 無名関数
	func(msg string) {
		fmt.Println(msg)
	}("John")

	// クロージャ
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}