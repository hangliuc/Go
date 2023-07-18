package main

import "fmt"

func test(arr [3]int) { //[3]int [4]int长度不一样，属于两个不同的类型
	arr[0] = 88
}

func test02(arr *[3]int) {
	(*arr)[0] = 88

}
func main() {
	//1、数组是相同类型的组合，一个数组一旦声明，其长度固定，不能动态变化
	var arr01 [3]int
	arr01[0] = 1
	arr01[1] = 2
	//arr01[2] = 1.1  数据类型一定要相同，否则会报错

	fmt.Println(arr01)

	//2、var arr []int 这时arr是一个切片

	//3、数组中的元素既可以是值类型也可以是引用类型

	/*4、数组地默认值
	数值类型 "0"
	字符串 ""
	bool false

	*/

	// 5、数组的下标是从0开始

	//6、数组的下标必须在指定范围内使用，否则报panic；数组越界

	//7、Go 数组属值类型，默认情况下是值传递、因此会进行值拷贝、数组间不会相互影响
	arr := [3]int{11, 22, 33}
	test(arr)
	fmt.Println(arr)

	//输出的还是 [11 22 33]

	// 8、如果想改变原来的的值

	test02(&arr)
	fmt.Println(arr)
	
	//长度是数组类型的一部分，在传递函数时，需要考虑数组的长度
}
