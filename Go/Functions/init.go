package main

import "fmt"

func init() {
	fmt.Println("init()...")

}

// init 函数会被main 函数之前被调用

// 通常可以在init 函数中完成初始化工作
func main() {
	fmt.Println("main()....")
	
}
