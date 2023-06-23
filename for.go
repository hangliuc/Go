package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ { // 打印层数
		for j := 1; j <= i; j++ { //打印列数
			fmt.Printf("%v * %v = %v ", j, i, j*i)
		}
		fmt.Println()
	}
}

// 函数调用
func printMuIti(num int) {
	for i := 1; i <= num; i++ { // 打印层
		for j := 1; j <= i; j++ { //打印列数
			fmt.Printf("%v * %v = %v ", j, i, j*i)
		}
		fmt.Println()
	}

}

func main() {
	var num int
	fmt.Println("请输入乘法表对应数")
	fmt.Scanln(&num)
	printMuIti(num)
}
