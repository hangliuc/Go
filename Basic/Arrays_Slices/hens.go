package main

import "fmt"

func main() {
	// go中数组是 值类型
	hen1 := 3.0
	hen2 := 5.0
	hen3 := 1.0
	hen4 := 3.4
	hen5 := 2.0
	hen6 := 50.0

	totalweight := hen1 + hen2 + hen3 + hen4 + hen5 + hen6
	avgweight := fmt.Sprintf("%.2f", totalweight/6) // totalweight/6 保留到小数点两位
	fmt.Printf("totalweight=%v avgweight=%v", totalweight, avgweight)

	// 使用数组解决
	var hens [6]float64
	// 初始化 给数组赋值
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0

	//遍历数组求总体重
	totalweight2 := 0.0
	for i := 0; i < len(hens); i++ {
		totalweight2 += hens[i]

	}

	avgweight2 := fmt.Sprintf("%.2f", totalweight2/float64(len(hens)))
	fmt.Printf("totalweight=%v avgweight=%v", totalweight2, avgweight2)

}
