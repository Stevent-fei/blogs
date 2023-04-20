package main

import "fmt"

/*
new 和 make的区别
make和new都是golang用来分配内存的內建函数，且在堆上分配内存，make 即分配内存，也初始化内存。new只是将内存清零，并没有初始化内存。
make返回的还是引用类型本身；而new返回的是指向类型的指针。
make只能用来分配及初始化类型为slice，map，channel的数据；new可以分配任意类型的数据
 */

func main() {

	var s []string
	a := []string{"1","2"}
	for _,b := range a{
		s = append(s,b)
	}
	//fmt.Println(s)

	var i int
	//s2 := string(i)
	fmt.Println(i)

	//make是


	//new是指针类型,能接收任何类型
	//list := new([]int)
	//list = append(list, 1)
	//fmt.Println(list)
}
