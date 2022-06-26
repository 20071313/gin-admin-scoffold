package main

import (
	"fmt"
	"github.com/spf13/cast"
	"time"
)

var msg = make(chan int)
var count = 0

func main() {
	go send()
	go receive()
	time.Sleep(1 * time.Minute)
}

func send() {
	for {
		msg <- count
		fmt.Println("向通道中导入count:" + cast.ToString(count))
		count++
		time.Sleep(2 * time.Second)
	}
}

func receive() {
	for {
		countStr := <-msg
		fmt.Println("从读取中读取:", countStr)
	}
}
