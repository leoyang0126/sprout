package main

import (
	"sprout/api"
	"sprout/app"
	"sprout/utils"
)

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	println("测试应用")
	testAclass.Test()

	var book = Book{"JAVA", "张三", "李思思", 12313}
	b := Book{}

	book1 := Book{"sdf", "李四", "sdf", 1231}
	println(book.book_id)
	println(b.subject)
	println(book1.book_id)

	//for i := 0; i < 10; i++ {
	//	testAclass.Login("张三", "sss")
	//}
	//utils.ArrayUnit()
	utils.Test()
	testAclass.TestStr()
	api.Start()
}
