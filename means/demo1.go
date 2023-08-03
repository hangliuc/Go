package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) speak() {
	fmt.Println(p.Name, "是一个goodman")

}
func (p Person) test() {
	fmt.Println("test() name=", p.Name)

}

func (p Person) jisuan() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i

	}
	fmt.Println("计算的结果是", res)

}

func (p Person) jisuan2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算的结果是", res)

}

func (p Person) getsum(n1 int, n2 int) int {

	return n1 + n2
}
func main() {
	var p Person

	p.Name = "tom"
	p.test()
	p.speak()
	p.jisuan()
	p.jisuan2(10)
	p.getsum(10, 20)

}
