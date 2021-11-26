package main

//生成随机数
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num1 := rand.Float64()
	fmt.Println(num1)

	for i := 0; i < 10; i++ {
		num := rand.Intn(10)
		fmt.Println(num)
	}
	rand.Seed(90)
	num2 := rand.Intn(10)
	fmt.Println("num2: ", num2)

	t1 := time.Now()
	fmt.Println(t1)
	fmt.Printf("%T", t1)

	fmt.Println("==========")
	//时间戳 秒，纳秒
	t2 := time.Now().Unix()
	fmt.Println(t2)

	t3 := time.Now().UnixNano()
	fmt.Println(t3)

	//step1，设置种子数seed，可以设置为不停变化的时间戳
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		//step2，调用生成随机数的函数
		fmt.Println("-->", rand.Intn(100))
	}

	/*
		获取指定范围的随机数
		[3,48]
		   [0,45]+3
		[15,76]
		   [0,61]+15
		rand.Intn(n)  --> [0,n)
	*/

	num3 := rand.Intn(46) + 3 //[3,48]
	fmt.Println(num3)

	num4 := rand.Intn(62) + 15
	fmt.Println(num4)
}
