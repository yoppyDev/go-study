package main

import (
	"fmt"
	"sync"
)

func chanGoroutine(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func chanOutPut(s []int, chanA chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		chanA <- sum
	}
	close(chanA)
}

func goroutine(wg *sync.WaitGroup) () {
	defer wg.Done()
	fmt.Println("hello")
}

func normal () {
	fmt.Println("world")
}

func main () {
	// WaitGroupを使って並列処理（goroutine）が終わるまで待つ
	var wg sync.WaitGroup
	wg.Add(1)
	goroutine(&wg)
	normal()

	// チャネルを使うと、goroutineの終了を待つことができる
	s := []int{1, 2, 3, 4, 5}
	g := []int{1, 2, 3}
	c := make(chan int)
	go chanGoroutine(s, c)
	go chanGoroutine(g, c)
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)

	// チャネルを使って、都度出力する
	chanA := make(chan int, len(s)) // 数がわかっていればメモリを確保
	go chanOutPut(s, chanA)
	for i := range chanA {
		fmt.Println(i)
	}
}