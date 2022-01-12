package main

import "fmt"

func main() {
	/*
	  打印2->100内所有的素数(只能被1和本身整除)
	*/
	for i := 2; i < 100; i++ {
		flag := true //记录i是否为素数
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false //不是素数
				break
			}
		}
		if flag == true {
			fmt.Printf("素数为：%d\n", i)
		}
	}
}
