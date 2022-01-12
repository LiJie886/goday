package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("已有数组直接创建切片-------------------------")
	s1 := a[:5]
	s2 := a[3:8]
	s3 := a[5:]
	s4 := a[:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

}
