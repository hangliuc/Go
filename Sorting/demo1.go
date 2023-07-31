package main

/*
内部排序
将需要处理的所有数据都加载到内部的存储器进行排序
交换式排序、选择式排序法和插入式排序法

外部排序
数据量过大，无法全部加载到内存中，需要借助外部存储进行排序
合并排序和直接合并排序法

*/

//交换排序---冒泡排序
import "fmt"

func BubbleSort(arr *[5]int) {

	fmt.Println("排序前arr=", (*arr))
	temp := 0 // 临时变量
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后arr=", (*arr))

	//fmt.Println("第一次排序arr=", (*arr))
	//
	//// 第二轮排序
	//for j := 0; j < 3; j++ {
	//	if (*arr)[j] > (*arr)[j+1] {
	//		temp = (*arr)[j]
	//		(*arr)[j] = (*arr)[j+1]
	//		(*arr)[j+1] = temp
	//	}
	//}
	//// 第三轮排序
	//fmt.Println("第二次排序arr=", (*arr))
	//
	//for j := 0; j < 2; j++ {
	//	if (*arr)[j] > (*arr)[j+1] {
	//		temp = (*arr)[j]
	//		(*arr)[j] = (*arr)[j+1]
	//		(*arr)[j+1] = temp
	//	}
	//}
	//
	//fmt.Println("第三次排序arr=", (*arr))
	//
	//// 第四轮排序
	//fmt.Println("第二次排序arr=", (*arr))
	//
	//for j := 0; j < 1; j++ {
	//	if (*arr)[j] > (*arr)[j+1] {
	//		temp = (*arr)[j]
	//		(*arr)[j] = (*arr)[j+1]
	//		(*arr)[j+1] = temp
	//	}
	//}
	//
	//fmt.Println("第四次排序arr=", (*arr))
}
func main() {
	arr := [5]int{24, 69, 80, 57, 13}
	BubbleSort(&arr)
}
