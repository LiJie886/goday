package main

import "fmt"

//const 定义枚举类型
const (
	//可以在const()中添加一个关键字iota，每行的iota都会累加，第一行的iota默认值为0
	BEIJING  = 10 * iota //iota=0
	SHANGHAI             //iota=1
	SHENZHEN             //iota=2
)

//iota只能配合const()一起使用
const (
	a, b = iota + 1, iota + 2 //iota=0,a=1,b=2
	c, d                      //iota=1,c=2,d=3
	e, f                      //iota=2,e=3,f=4

	g, h = iota * 2, iota * 3 //iota=3,g=6,h=9
	i, k                      //iota=4,i=8,k=12
)

func main() {
	const length = 10 //常量只读，不可被修改
	fmt.Println("length = ", length)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("Shanghai = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)
	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "k = ", k)
}
