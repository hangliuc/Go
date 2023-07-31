package main

import "fmt"

func main() {
	// 声明不会分配内存，初始化需要make，分配内存后才能使用

	//第一种方式
	var a map[string]string

	a = make(map[string]string, 10)

	a["no1"] = "宋江"
	a["no2"] = "林冲"
	a["no3"] = "李逵"
	a["no4"] = "武松"
	fmt.Println(a)
	//第二种方式

	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "西安"
	fmt.Println(cities)

	//第三种方式
	heroes := map[string]string{
		"hero1": "tom",
		"hero2": "jack",
		"hero3": "hang",
	}
	heroes["hero4"] = "feifei"

	fmt.Println(heroes)

}
