package main

import (
	"fmt"
	"sync"
)

// NumberPrint beginからendまでの値を表示する
func NumberPrint(begin int, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := begin; i <= end; i++ {
		fmt.Println(i)
	}
}

// NumberAddition int型スライスの合計値を1計算ごとにチャネルに返す
func NumberAddition(numbers []int, c chan int) {
	sum := 0
	defer close(c)
	for _, value := range numbers {
		sum += value
		// チャネルへ値を送る
		c <- sum
	}
}

func main() {
	// 2つのNumberPrint関数を並行処理
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go NumberPrint(1, 10, wg)
	go NumberPrint(100, 110, wg)
	wg.Wait()

	// チャネルを使ってgoroutineの戻り値を受け取る
	numbers := []int{1, 3, 5, 7, 9}
	c := make(chan int)
	go NumberAddition(numbers, c)
	// チャネルの要素を拾う
	for i := range c {
		fmt.Println(i)
	}
}
