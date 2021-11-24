package main

import "fmt"

/*
四种变量的声明方式
*/
func main() {
	// 方法一，声明一个变量,默认值为0
	var a int
	fmt.Println("a =", a)
	fmt.Printf("type of a =%T\n", a)
	//  声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b =%T\n", b)
	//在初始化时，可以省去数据类型，通过值自动匹配数据类型
	var c = 20
	fmt.Println("c =", c)
	fmt.Printf("type of c =%T\n", c)
	//方法四，省去var关键字，直接自动匹配
	d := 0.0
	fmt.Printf("type of d =%T", d)

	fmt.Println()

	var e, f = 100, 200
	fmt.Println("e = ", e, "f = ", f)
	var (
		g = 100
		h = "haode"
	)
	fmt.Println("g = ", g)
	fmt.Println("h =", h)
}
