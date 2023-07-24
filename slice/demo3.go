package main

import "fmt"

func main() {
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	fmt.Println(len(slice))
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	}

	fmt.Println()

	for i, v := range slice { // i数组的下标，数组的值
		fmt.Printf("i=%v v=%v\n", i, v)
	}
}
