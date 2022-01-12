package main

import (
	"fmt"
)

func main() {
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5] //[2]->[4]=[90,82,100]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++ //[2]->[4]=[91,83,101]
	}
	fmt.Println("array after", darr)
}
