package main

import "fmt"

type person struct {
	Name string
}

func test01(p Person) {
	fmt.Println(p.Name)
}

func test02(p *Person) {
	fmt.Println(p.Name)
}

func main() {

	p := Person{"tom"}
	test01(p)
	test02(&p)
}
