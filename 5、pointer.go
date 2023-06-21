// package main
//
// import "fmt"
//
//	func main() {
//		//基本数据类型在内存布局
//		//指针变量存的是一个地址，这个地址指向的空间才是值
//		var i int = 10
//		// 取出i的地址，&i
//		fmt.Println("i的地址：", &i)
//
//		// 1、ptr是一个指针变量
//		// 2、ptr 的类型 *int
//		// 3、ptr 本身的值是&i
//		var ptr *int = &i
//		fmt.Printf("ptr=%v\n", ptr)
//		fmt.Printf("ptr的地址=%v\n", &ptr)
//		fmt.Printf("ptr 指向的值=%v\n", *ptr)
//	}
package main

import "fmt"

func main() {

	var num int = 9
	fmt.Println("num 的地址=", &num)

	var ptr *int
	ptr = &num
	*ptr = 10 // 这里修改时会导致 num值变化
	fmt.Println("num = ", num)

}
