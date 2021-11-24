package main

import "fmt"

//值传递
/* func swap(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
} */
func swap(pa *int, pb *int) {
	var temp int
	temp = *pa //temp = main::a
	*pa = *pb  //main::a = main::b
	*pb = temp //main::b = temp
}

func main() {
	var a int = 10
	var b int = 20
	swap(&a, &b)
	fmt.Println("a = ", a, "b = ", b)
}
