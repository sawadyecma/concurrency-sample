package main

import (
	"fmt"
	"sync"
)

func noGoroutine() {
	for i := 0; i < 5; i++ {
		fmt.Println("No goroutine")
	}
}

// WaitGroupのアドレスを渡す
func goroutine(wg *sync.WaitGroup) {
	// WaitGroupにタスク完了を通知
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("Yes goroutine")
	}
}

func main() {
	// WaitGroup作成
	var wg sync.WaitGroup
	// スレッドが一つあることを伝える
	wg.Add(1)
	go goroutine(&wg)
	noGoroutine()
	// スレッド完了まで待機
	wg.Wait()
}
