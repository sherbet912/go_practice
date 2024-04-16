package main

import (
	. "fmt"
)

type Human struct {
	name string
	age int
	weight float64
	phone string
}

type Student struct {
	Human // 匿名欄位, 預設 Student 包含 Human 所有欄位
	speciality string
	phone string
}

func main() {
	mark := Student{Human{name: "Mark", age: 15, weight: 66.6, phone: "111111111"}, "CS", "999999999"}

	Println("A student")
	Println("His name is:", mark.name)
	Println("His age is:", mark.age)
	Println("His weight is:", mark.weight)
	Println("His speciality is:", mark.speciality)

	mark.speciality = "EE"
	Println("Mark changed his speciality:")
	Println("His speciality is:", mark.speciality)

	Println("Mark become old")
	mark.age = 30
	Println("His age is:", mark.age)

	Println(mark.Human)
	// 如果有相同欄位, 則外層會先存取, 內層要指定進去 (類似於繼承)
	Println(mark.phone, mark.Human.phone)
}