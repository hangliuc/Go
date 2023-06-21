// 一个指针变量指向了一个值的内存地址。
// 指针的声明格式 var var_name *var-type
package main

import "fmt"

func main() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */ //取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址

	fmt.Printf("a 变量的地址是: %v\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %v\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %v\n", *ip) //在指针类型前面加上 * 号（前缀）来获取指针所指向的内容
}
