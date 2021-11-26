package main

import "fmt"

func main() {
	s1 := make([]int, 0, 4)
	s2 := make([]int, 10, 99)
	fmt.Println(s1)
	s1 = append(s1, 1, 2, 3, 4, 5, 6)
	s2 = append(s2, 44, 55, 66, 77, 88)
	fmt.Println(s1)
	s1 = append(s1, s2...)
	fmt.Println(s1)

	for i := 0; i < len(s1); i++ {
		fmt.Println(s1)
	}
	for i, v := range s1 {
		fmt.Printf("%d-->%d\n", i, v)
	}
}
