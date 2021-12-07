package main

import "fmt"

//递归函数
func main() {
	fmt.Println(getSum(5))
}
func getSum(n int) int {
	//1.求1-->5的和
	if n == 1 {
		return 1
	}
	return getSum(n-1) + n
}
