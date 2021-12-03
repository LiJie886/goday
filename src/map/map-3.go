package main

import (
	"fmt"
	"sort"
)

/*
map遍历
使用for range

数组，切片：index,value,
map:key,value
*/
func main() {
	map1 := make(map[int]string)
	map1[1] = "孙悟空"
	map1[2] = "猪八戒"
	map1[3] = "沙和尚"
	map1[4] = "唐僧"
	map1[5] = "哪吒"
	map1[6] = "红孩儿"

	//1.遍历map
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	fmt.Println("----------------------")
	for i := 1; i <= len(map1); i++ {
		fmt.Println(i, "-->", map1[i])

	}
	/*
	   1.获取key
	   2.排序
	   3.遍历map
	*/
	keys := make([]int, 0, len(map1))
	fmt.Println(keys)

	for k, _ := range map1 { //只要key不要value，用_表示
		keys = append(keys, k)
	}
	fmt.Println(keys)

	sort.Ints(keys) //默认升序
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key, map1[key])
	}

	s1 := []string{"apple", "orange", "banana", "windows", "linux", "API", "张三", "李四", "王五", "张六", "章草"}
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)
}
