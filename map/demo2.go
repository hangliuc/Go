package main

import "fmt"

func main() {
	//map[string]map[string]string

	StudentMap := make(map[string]map[string]string)

	StudentMap["stu01"] = make(map[string]string, 3)
	StudentMap["stu01"]["name"] = "tom"
	StudentMap["stu01"]["sex"] = "男"
	StudentMap["stu01"]["address"] = "长安街"

	StudentMap["stu02"] = make(map[string]string, 3)
	StudentMap["stu02"]["name"] = "mary"
	StudentMap["stu02"]["sex"] = "女"
	StudentMap["stu02"]["address"] = "沙河"

	fmt.Println(StudentMap)
	fmt.Println(StudentMap["stu02"])
	fmt.Println(StudentMap["stu02"]["sex"])
}
