package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello北" // golang 统一编码为utf-8 （ascii 的 字母和数字占一个字符，汉字占用三个字符// ）
	fmt.Println("str len=", len(str))

	str2 := "hello北京"
	// 2、字符串遍历，同时有中文问题 r := []rune(str) 转换rune的切片
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	//3、字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换的结果是", n)
	}

	//4、整数转字符串
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T\"", str, str)
	println()

	//5, 字符串转 []byte
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	//6、[]byte 转字符串

	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

	//7、10 进制 转 2，8，16
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是=%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的十六进制是=%v\n", str)

	//8、查找子串是否在指定的字符串中
	b := strings.Contains("seafoof", "oo")
	fmt.Println(b)

	// 9、统计一个字符串有几个子串
	num := strings.Count("seafoof", "f")
	fmt.Println(num)

	//10、不区分大小写（ == 区分大小写 ）
	b = strings.EqualFold("ABC", "abc")
	fmt.Printf("b=%v\n", b)

	fmt.Println("结果", "ABC" == "abc")
	
	//11、返回子串第一次出现index值
	index := strings.Index("NLT_abc", "abc")
	fmt.Printf("index=%v\n", index)
}
