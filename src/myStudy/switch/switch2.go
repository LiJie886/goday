package main

import (
	"fmt"
)

func main() {
	switch x := 5; x {

	case 5:
		x += 10
		fmt.Println(x)
		fallthrough
	case 6:
		x += 20
		fmt.Println(x)
	default:
		fmt.Println(x)
	}

}
