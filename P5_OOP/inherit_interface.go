package main

import (
	. "fmt"
	"strconv"
)

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

func (h Human) String() string {
	return "< " + h.name + " - " + strconv.Itoa(h.age) + " years - phone: " + h.phone + " >"
}

func (h Human) SayHi() {
	Printf("Hi, I am %s. You can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	Println("La~, la la la la, la...", lyrics)
}

// overwrite employee method SayHi()
func (e Employee) SayHi() {
	Printf("Hi, I am %s. I work at %s. Can call me on %s\n", e.name, e.company, e.phone)
}

func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}

func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

type YoungChap interface {
    SayHi()
    Sing(song string)
    BorrowMoney(amount float32)
}

type ElderlyGent interface {
    SayHi()
    Sing(song string)
    SpendSalary(amount float32)
}

type Element interface{}
    type List [] Element

type Person struct {
	name string
	age int
}

//定義了 String 方法，實現了 fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"
}

func main() {
	carlor := Student{Human{name: "Carlor", age:18, phone: "0977123321"}, "NTU", 0}
	kate := Employee{Human{name: "Kate", age: 25, phone: "0965432321"}, "TSMC", 5555}
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
    paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
    sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
    tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	carlor.SayHi()
	kate.SayHi()

	var i Men
	i = mike
	Println("This is Mike, a student.")
	i.SayHi()
	i.Sing("Hello 你好嗎～")

	i = tom
	Println("This is Tom, an employee")
	i.SayHi()
	i.Sing("Believe it~")

	Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, kate

	for _, value := range x {
		value.SayHi()
	}

	Println("This human is:", carlor)

	list := make(List, 3)
	list[0] = 1 // an int
	list[1] = "Hello" // a string
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		if value, ok := element.(int); ok {
			Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			Printf("list[%d] is of a different type\n", index)
		}

		switch value := element.(type) {
			case int:
				Printf("list[%d] is an int and its value is %d\n", index, value)
			case string:
				Printf("list[%d] is a string and its value is %s\n", index, value)
			case Person:
				Printf("list[%d] is a Person and its value is %s\n", index, value)
			default:
				Printf("list[%d] is of a different type\n", index)
		}
	}
}