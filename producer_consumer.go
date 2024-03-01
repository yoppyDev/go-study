package main

import (
	"fmt"
	"sync"
)

/**
	producer　と　consumerについて
	producer -> channel -> consumer
	producerは、channelにデータを送信する
	consumerは、channelからデータを受信する
	consumerはrangeでchannelからデータを受信する
**/

func producer(c chan string, v string) {
	c <- v
}


func consumer(c chan string, wg *sync.WaitGroup) {
	for v := range c {
		func () {
			// wg.Done()で、goroutineの終了を通知する
			defer wg.Done()
			fmt.Println("Received", v)
		}()
	}
}


func main() {
	// sync.WaitGroupを使って、goroutineの終了を待つ
	var wg sync.WaitGroup
	ary := []string{"a", "b", "c", "d", "e"}
	c := make(chan string)

	for _, v := range ary {
		// wg.Add(1)で、goroutineの数をカウントする
		wg.Add(1)
		go producer(c, v)
	}

	go consumer(c, &wg)
	// wg.Wait()で、goroutineの終了を待つ
	wg.Wait()
	// channelをcloseする
	close(c)
	fmt.Println("Done")
}
