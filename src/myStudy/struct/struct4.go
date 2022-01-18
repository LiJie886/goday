package main

import "fmt"

func main() {

	/*
		结构体嵌套：一个结构体中的字段，是另一个结构体类型。
	*/
	b1 := Book{}
	b1.bookName = "西游记"
	b1.price = 18.8

	s1 := Student1{}
	s1.name = "zzhsngsan"
	s1.age = 20
	s1.book = b1 //值传递
	fmt.Println(b1)
	fmt.Println(s1)
	fmt.Printf("学生姓名：%s，学生年龄：%d，看的书是：《%s》，书的价格是:%.2f\n", s1.name, s1.age, s1.book.bookName, s1.book.price)

	s1.book.bookName = "红楼梦"
	fmt.Println(s1)
	fmt.Println(b1)

	s2 := Student1{"xiaohua", 30, Book{"镜花缘", 23.66}}
	fmt.Printf("学生姓名：%s，学生年龄：%d，看的书是：《%s》，书的价格是:%.2f\n", s2.name, s2.age, s2.book.bookName, s2.book.price)
	s3 := Student1{
		name: "李华",
		age:  20,
		book: Book{
			bookName: "未解之谜",
			price:    30.06,
		},
	}
	fmt.Printf("学生姓名：%s，学生年龄：%d，看的书是：《%s》，书的价格是:%.2f\n", s3.name, s3.age, s3.book.bookName, s3.book.price)
	fmt.Println(s3.name, s3.age)
	fmt.Println("\t", s3.book.bookName, s3.book.price)

	b4 := Book{"三国志", 88.88}
	s4 := Student2{"韩梅梅", 24, &b4}
	fmt.Println(s4)
	fmt.Println("\t", s4.book)

	s4.book.bookName = "水浒传"
	fmt.Println(b4)
	fmt.Println(s4.book)

}

//1.定义一本书的结构体
type Book struct {
	bookName string
	price    float64
}

//2.定义学生的结构体
type Student1 struct {
	name string
	age  int
	book Book
}
type Student2 struct {
	name string
	age  int
	book *Book //Book的地址

}
