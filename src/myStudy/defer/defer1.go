package main

import "fmt"

//defer在return之后执行
func deferFunc() int {
	fmt.Println("defer.......")
	return 0
}
func returnFunc() int {
	fmt.Println("return......")
	return 0
}
func deferAndReturn() int {
	defer deferFunc()
	return returnFunc()
}
func main() {
	deferAndReturn()
}
