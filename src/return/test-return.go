package main

import "fmt"

//返回单个返回值
func fool(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

//匿名返回多个返回值
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 666, 777
}

func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("----foo3----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("----foo4----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return

}
func main() {
	c := fool("abc", 555)
	fmt.Println("c = ", c)

	re1, re2 := foo2("ha", 999)
	fmt.Println("re1 = ", re1, "re2 = ", re2)

	re3, re4 := foo3("foo3", 444)
	fmt.Println("re3 = ", re3, "re4 = ", re4)

	re5, re6 := foo4("foo4", 111)
	fmt.Println("re5 = ", re5, "re6 = ", re6)

}
