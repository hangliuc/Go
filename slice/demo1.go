package main

import "fmt"

func main() {
	// 不确定个数的数组，长度可以变化，是一个动态变化的数组
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}

	//slice 起始下标为1 ，最后下标为3（但是不包含3）
	slice := intArr[1:3]

	fmt.Println("intArr", intArr)
	fmt.Println("slice的元素是", slice)
	fmt.Println("slice的容量是", cap(slice))

}
