package main

import "fmt"

func main() {

	var intArr [3]int //int 占8个字节
	//当我们定义完数组，其实数组的各个元素是0
	//1、数组的地址可以&intAr得出 
	//2、数组的第一个元素地址，就是数组的首地址
	//3、数组的各元素的地址间隔，根据数组类型决定 int64 = 8 int32 =4
	fmt.Println(intArr)

	fmt.Printf("intArr的地址=%p intArr[0] 地址是%p intArr[1] 地址是%p intArr[2] 地址是%p ",
		&intArr, &intArr[0], &intArr[1], &intArr[2]) // %p 取出地址
}
