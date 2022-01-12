package main

import "fmt"

func main() {
	s1 := "hello world"
	s2 := "你好呀"
	s3 := `good`
	fmt.Println(s1, "\t", s2, "\t", s3)
	fmt.Println(len(s1), len(s2), len(s3))

	//字符串是字节的集合
	slice1 := []byte{65, 66, 67, 68, 69}
	s4 := string(slice1) //根据字节切片构建字符串
	fmt.Println(s4)

	s5 := "abcde"        //字符串不能修改
	slice2 := []byte(s5) //根据字符串，获取对应的字符切片
	fmt.Println(slice2)

}
