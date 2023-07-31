package main

import "fmt"

func main() {
	names := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	// 顺序查找
	var heroName = ""
	fmt.Println("请输入要查找的人名")
	fmt.Scanln(&heroName)

	//for i := 0; i < len(names); i++ {
	//	if heroName == names[i] {
	//		fmt.Printf("找到%v,下标%v", heroName, i)
	//		break
	//	} else if i == (len(names) - 1) {
	//		fmt.Printf("没有找到%v", heroName)
	//	}
	//
	//}

	//顺序查找
	index := -1
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			index = i
			break
		}
	}

	if index != -1 {
		fmt.Printf("找到%v,下标%v", heroName, index)
	} else {
		fmt.Println("没有找到")
	}

}
