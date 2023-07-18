package main

import "fmt"

func main() {
	//
	//var score [5]float64
	//
	//for i := 0; i < len(score); i++ {
	//	fmt.Println("请输入第%d个元素的值", i+1)
	//	fmt.Scanln(&score[i])
	//}
	//
	//// 数组的打印
	//
	//for i := 0; i < len(score); i++ {
	//	fmt.Printf("score[%d]=%v\n", i, score[i])
	//}

	//4种初始化数组的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01=", numArr01)

	var numArr02 = [3]int{5, 6, 7}
	fmt.Println("numArr01=", numArr02)

	var numArr03 = [...]int{8, 9, 10}
	fmt.Println("numArr01=", numArr03)

	var numArr04 = [...]int{1: 800, 0: 900, 2: 1000}
	fmt.Println("numArr01=", numArr04)

}
