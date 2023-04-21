package main

import (
	"fmt"
	"kubernetes/staging/src/k8s.io/apimachinery/pkg/util/json"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	High  int    `json:"high"`
	Sex   string `json:"sex"`
	Class *Class `json:"class"`
}

type Class struct {
	ClassName string `json:"classname"`
	Grade     int    `json:"grade"`
}

func main() {

	person := Student{
		Name: "张三",
		Age:  18,
		High: 182,
		Sex:  "男",
	}

	//指针变量
	class := new(Class)
	class.ClassName = "1班"
	class.Grade = 3
	person.Class = class

	jsonbyte, err := json.Marshal(person)
	if err != nil {
		return
	}
	fmt.Println(string(jsonbyte))

	stu := Student{}
	err = json.Unmarshal(jsonbyte, &stu)
	//解析失败会报错，如json字符串格式不对，缺"号，缺}等。
	if err != nil {
		return
	}
	fmt.Println(stu)
}
