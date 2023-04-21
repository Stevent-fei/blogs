package main

import "fmt"

func main() {
	var intArr = [...]int{1, 0, 2, 3} //test02
	fmt.Println(len(intArr))
	for i := 0; i < len(intArr)-1; i++ {
		for j := 0; j < len(intArr)-1-i; j++ {
			if intArr[j+1] < intArr[j] {
				temp := intArr[j+1]
				intArr[j+1] = intArr[j]
				intArr[j] = temp
			}
		}
		fmt.Printf("第%v轮冒泡排序后：%v\n", i+1, intArr)
	}
}
