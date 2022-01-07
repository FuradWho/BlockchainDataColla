package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func myPrintA() {
	defer wg.Done()
	fmt.Println("A")
}
func myPrintB() {
	defer wg.Done()
	runtime.Gosched() //打印B之前，让出当前goroutine所占的时间片
	fmt.Println("B")
}
func main() {
	runtime.GOMAXPROCS(1)
	for i := 1; i <= 10; i++ {
		wg.Add(2)
		go myPrintA()
		go myPrintB()
		time.Sleep(time.Second)
	}
	wg.Wait()
}
