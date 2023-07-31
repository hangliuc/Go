package main

import "fmt"

func main() {
	var arr [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr3=", arr)

	// for 遍历

	for i := 0; i < len(arr); i++ { // 0
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v\t", arr[i][j])
		}
		fmt.Println()
	}
}
