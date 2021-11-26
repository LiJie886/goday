package main

import "fmt"

//水仙花数：三位数
//每个位上数字的立方和加起来等于该数本身
func main() {
	//方法一
	for i := 100; i < 999; i++ { //153
		ge := i % 10       //ge == 3
		shi := i / 10 % 10 //shi == 5
		bai := i / 100     //bai == 1
		if i == ge*ge*ge+shi*shi*shi+bai*bai*bai {
			fmt.Printf("水仙花数为：%d\n", i)
		}
	}

	fmt.Println("-------------------------")
	//方法二
	for a := 1; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				n := a*100 + b*10 + c*1
				if a*a*a+b*b*b+c*c*c == n {
					fmt.Printf("水仙花数为：%d\n", n)
				}
			}
		}
	}
}
