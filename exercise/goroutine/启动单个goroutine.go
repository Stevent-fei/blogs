package main

import (
	"fmt"
	"time"
)

func hi() {
	fmt.Println("hello goroutine")
}

func main() {
	go hi()
	fmt.Println("main goroutine done")
	time.Sleep(time.Second) //必须等待，不然执行完打印就会结束，原因是因为创建goroutine需要花费一些时间
}