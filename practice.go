package main

import "fmt"

func main() {
	var arr [5]int = [...]int{-1, 1, 10, 20, 30}
	sum := 0
	sum1 := 0
	for _, val := range arr {
		sum += val
	}

	fmt.Println(sum)

	for i := 0; i < len(arr); i++ {
		sum1 += arr[i]
	}

	fmt.Println(sum1)
}
