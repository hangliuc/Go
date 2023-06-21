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
