package main

import (
	"fmt"
	"time"
)

/*
	select文は、複数のチャネルを待ち受けることができる。
	queueのように、複数のgoroutineからのデータを受け取る場合に有効。
*/

func goroutineSendMail(ch chan string) {
	fmt.Println("send mail" + <-ch)
}

func goroutinePacket1(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(100 * time.Millisecond)
	}
}

func goroutinePacket2(ch chan string) {
	for {
		ch <- "packet from 2"
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	boom := time.After(500 * time.Millisecond)
	go goroutinePacket1(c1)
	go goroutinePacket2(c2)


	loop:
		for {
			select {
			case <-c1:
				go goroutineSendMail(c1)
			case <-c2:
				go goroutineSendMail(c2)
			case <-boom:
				fmt.Println("BOOM!")
				break loop
			default:
				fmt.Println("...")
				time.Sleep(50 * time.Millisecond)
			}
		}
}
