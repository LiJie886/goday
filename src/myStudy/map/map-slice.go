package main

import "fmt"

func main() {
	/*
		map与slice的结合使用
		1.创建map用于储存人的信息
		   name，age，sex，address
		2.每个map存储一个人的信息
		3.讲这些map存入slice中
		4.打印遍历输出
	*/

	//1.map创建第一个人信息
	map1 := make(map[string]string)
	map1["name"] = "张三"
	map1["age"] = "30"
	map1["sex"] = "男"
	map1["address"] = "上海"
	fmt.Println(map1)

	//2.第二个人的信息
	map2 := make(map[string]string)
	map2["name"] = "李四"
	map2["age"] = "28"
	map2["sex"] = "女"
	map2["address"] = "北京"
	fmt.Println(map2)

	map3 := map[string]string{"name": "小花", "age": "8", "sex": "母", "address": "狗窝"}
	fmt.Println(map3)

	s1 := make([]map[string]string, 0, 3)
	s1 = append(s1, map1)
	s1 = append(s1, map2)
	s1 = append(s1, map3)
	fmt.Println(s1)

	//遍历
	for i, val := range s1 {
		fmt.Printf("第 %d 个人的信息是: \n", i+1)
		fmt.Printf("\t姓名：%s\n", val["name"])
		fmt.Printf("\t年龄：%s\n", val["age"])
		fmt.Printf("\t性别：%s\n", val["sex"])
		fmt.Printf("\t地址：%s\n", val["address"])
	}
}
