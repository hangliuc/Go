package main

import (
	"fmt"
	"strings"
)

// 闭包 一个函数与其相关的引用环境组合的一个整体（实体）

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n

	}

}

/*
闭包：
	var n int = 10 //不会每次都初始化
	return func(x int) int {
		n = n + x
		return n
*/
// 1) AddUpper 是一个函数，返回的数据类型是func(int) int
// 2） 返回的是一个匿名函数func(x int) int，但是这个匿名函数引用到函数外的n，这个匿名函数和n 形成一个闭包
// 3）当我们反复调用f函数时，n是初始化一次，每调用一个就进行累计

func makeSuffix(suffix string) func(string) string {

	return func(name string) string {
		// 如果 name 没有指定的后缀名，则加上，否则按照原来的名字返回
		if !strings.HasSuffix(name, suffix) { // 等同于 strings.HasSuffix(name, suffix) == false
			return name + suffix
		}
		return name

	}

}

// 不使用闭包 
// 闭包只传递一次，函数得每次传递

func makeSuffix2(suffix string, name string) string {
	// 如果 name 没有指定的后缀名，则加上，否则按照原来的名字返回
	if !strings.HasSuffix(name, suffix) { // 等同于 strings.HasSuffix(name, suffix) == false
		return name + suffix
	}
	return name

}
func main() {
	// 累加器

	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	f2 := makeSuffix(".jpg")
	fmt.Println("文件处理后=", f2("winter"))

	fmt.Println("文件处理后=", makeSuffix2(".jpg", "winter"))
	fmt.Println("文件处理后=", makeSuffix2(".jpg", "bird.jpg"))

}
