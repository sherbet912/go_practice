package main

import (
	. "fmt" // 用 . import 可以省略掉 fmt. ex fmt.Println() > Println()
)

type person struct {
	name string
	age int
}

func older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}

	return p2, p2.age - p1.age
}

func main() {
	var p person
	p.name = "John"
	p.age = 20
	Printf("The person name is %s.\n", p.name)

	p2 := person{name: "Tom", age: 25}
	older_person, older_age := older(p, p2)
	Printf("%s and %s, %s is order by %d age", p.name, p2.name, older_person.name, older_age)
}