// 数组的定义
package main

import "fmt"

func main() {
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 30
	fmt.Println(slice)
	fmt.Println("slice的size=", len(slice))
	fmt.Println("slice的容量=", cap(slice))
	fmt.Println()
	var slice1 []string = []string{"tom", "jack", "mary"}
	fmt.Println(slice1)
	fmt.Println("slice1的size=", len(slice1))
	fmt.Println("slice1的容量=", cap(slice1))

}
