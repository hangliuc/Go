package main

import "fmt"

func main() {
	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i)
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", myChars[i])
	}

	fmt.Println()
	var intArr [5]int = [...]int{1, -1, 9, 90, 11}
	maxVal := intArr[0]
	maxValIndex := 0

	for i := 1; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxVal=%v maxValIndex=%v\n\n", maxVal, maxValIndex)

	var intArr2 [5]int = [...]int{1, -1, 9, 90, 12}
	sum := 0
	for _, val := range intArr2 {
		// 累计求和
		sum += val
	}
	fmt.Printf("sum=%v 平均值=%v", sum, float64(sum)/float64(len(intArr2)))
}
