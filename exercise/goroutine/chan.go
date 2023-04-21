package main

//func main() {
//
//	ch:=make(chan int,0)
//
//	go func(){
//		time.Sleep(time.Second*3)
//		fmt.Println("receive over")
//		a:=<-ch
//		fmt.Println(a)
//	}()
//
//	ch<-1
//	fmt.Println("send over")
//}

//import "fmt"
//
//func main() {
//	c := make(chan int)
//	go func() {
//		for i := 0; i < 5; i++ {
//			c <- i
//		}
//		close(c)
//	}()
//	for {
//		if data, ok := <-c; ok {
//			fmt.Println(data)
//		} else {
//			break
//		}
//	}
//	fmt.Println("main结束")
//}

import (
	"fmt"
	"time"
)

// 判断管道有没有存满
func main() {
	// 创建管道
	output1 := make(chan string, 10)
	// 子协程写数据
	go write(output1)
	// 取数据
	for s := range output1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}

func write(ch chan string) {
	for {
		select {
		// 写数据
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
