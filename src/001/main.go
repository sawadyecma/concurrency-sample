package main

import (
	"fmt"
	"time"
)

func noGoroutine() {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("No goroutine")
	}
}

func goroutine() {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Yes goroutine")
	}
}

func main() {
	go goroutine()
	noGoroutine()
}
