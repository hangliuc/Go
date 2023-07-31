package main

import "fmt"

func main() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "西安"
	fmt.Println(cities)

	for k, v := range cities {
		fmt.Printf("k=%v ,v=%v\n", k, v)

	}

	// 遍历复杂map
	StudentMap := make(map[string]map[string]string)
	StudentMap["stu01"] = make(map[string]string, 3)
	StudentMap["stu01"]["name"] = "tom"
	StudentMap["stu01"]["sex"] = "男"
	StudentMap["stu01"]["address"] = "长安街"

	StudentMap["stu02"] = make(map[string]string, 3)
	StudentMap["stu02"]["name"] = "mary"
	StudentMap["stu02"]["sex"] = "女"
	StudentMap["stu02"]["address"] = "沙河"

	for k1, v1 := range StudentMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v\n", k2, v2)
		}
	}
}
