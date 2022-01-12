package main

import "fmt"

func fun1() {
	fmt.Println("A")
}
func fun2() {
	fmt.Println("B")
}
func fun3() {
	fmt.Println("C")
}
func main() {
	//写入defer关键词
	defer fun1()
	defer fun2()
	defer fun3()
}
