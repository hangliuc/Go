package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func main() {
	// 方式一

	var person Person
	person.Name = "tom"
	person.age = 18

	fmt.Println(person)

	// 方式二

	p1 := Person{"mary", 20}

	fmt.Println(p1)

	//方式三

	var p2 *Person = new(Person)
	(*p2).Name = "joy"
	p2.Name = "fei"

	(*p2).age = 21
	p2.age = 25

	fmt.Println(*p2)

	//方式四

	var p3 *Person = &Person{}
	(*p3).Name = "xiaohong"
	p3.Name = "xiaoming"

	(*p3).age = 31
	p3.age = 30

	fmt.Println(*p3)
}
