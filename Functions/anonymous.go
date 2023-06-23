package main

import (
	"fmt"
)

// 匿名函数 某个函数只希望使用一次，匿名函数也可以实现多次调用

var (
	// 3、fun1 全局匿名函数
	fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	// 求两个数的和 使用匿名函数
	// 1、这种方式只能调用一次
	// 未设置 函数名 定义直接调用
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res1=", res1)

	// 将匿名函数 赋给 a 变量 可以重复调用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}

	res2 := a(10, 30)
	fmt.Println("res2=", res2)

	res3 := a(10, 30)
	fmt.Println("res3=", res3)

	// 3、全局匿名函数的调用
	res4 := fun1(4, 9)
	fmt.Println("res4=", res4)
}
