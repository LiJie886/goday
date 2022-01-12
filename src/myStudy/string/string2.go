package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	   string包下关于字符串的函数
	*/
	s1 := "helloworld"
	//是否包含指定内容-->bool
	fmt.Println(strings.Contains(s1, "abc"))
	//是否包含chars中任意一个字符-->bool
	fmt.Println(strings.ContainsAny(s1, "abcd"))
	//出现了几次
	fmt.Println(strings.Count(s1, "l"))

	//以xxx开头，以xxx结尾
	s2 := "20211202日"
	if strings.HasPrefix(s2, "2021") {
		fmt.Println("2021年")
	}
	if strings.HasSuffix(s2, "日") {
		fmt.Println("是今天")
	}

	fmt.Println(strings.Index(s2, "9"))

}
