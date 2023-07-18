package main

import "fmt"

func main() {
	//两种遍历数组方式
	heroes := [...]string{"宋江", "吴用", "卢俊义"}
	//1、for循环
	for x := 0; x < len(heroes); x++ {
		fmt.Printf("元素的值为=%v\n", heroes[x])
	}
	//2、for-range

	for i, v := range heroes {
		fmt.Printf("i=%v v=%v\n", i, v)
		fmt.Printf("heroes[%d]=%v\n", i, heroes[i])
	}

	for _, v := range heroes {
		fmt.Printf("元素的值为=%v\n", v)
	}
}
