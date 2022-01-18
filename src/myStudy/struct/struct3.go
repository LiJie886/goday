package main

import "fmt"

func main() {
	/*
		匿名结构体和匿名字段：
		匿名结构体：没有名字的结构体，
			在创建匿名结构体时，同时创建对象
			变量名 := struct{
				定义字段Field
			}{
				字段进行赋值
			}
		匿名字段：一个结构体的字段没有字段名
		匿名函数：
	*/

	s1 := Student{"zhangsan", 20}
	fmt.Println(s1.name, s1.age)

	func() {
		fmt.Println("hello!")
	}()

	s2 := struct {
		name string
		age  int
	}{
		"lisi",
		19,
	}
	fmt.Println(s2.name, s2.age)

	//w1:=Worker{"XIAOMO",22}
	//fmt.Println(w1.name,w1.age)

	w2 := Worker{"huawei", 44}
	fmt.Println(w2)
	fmt.Println(w2.string, w2.int)
}

type Worker struct {
	//name string
	//age int
	string //匿名字段
	int    //匿名字段，默认使用数据类型作为名字，那么匿名字段的类型就不能重复，否则会冲突

}

type Student struct {
	name string
	age  int
}
