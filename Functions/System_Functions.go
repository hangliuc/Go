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

	//12、返回最后出现的位置
	index = strings.LastIndex("go golang", "go")
	fmt.Printf("index=%v\n", index)

	//13、将指定子串替换另外的子串
	//n = -1 全部替换
	str = strings.Replace("go go hello", "go", "北京", -1)
	fmt.Printf("index=%v\n", str)

	//14、按照指定字符为分割符，将一个字符串拆分为字符串数组
	strArr := strings.Split("hello,world,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}
	fmt.Printf("strArr=%v\n", strArr)

	//15、将字符串进行大小写转换
	str = "goLang Hello"
	str = strings.ToLower(str)
	str = strings.ToUpper(str) // 大写
	fmt.Printf("str=%v\n", str)

	//16、将字符串左右空格去掉
	str = strings.TrimSpace("  i love you china   ")
	fmt.Printf("str=%v\n", str)

	//17、将字符串两边指定字符串去掉
	str = strings.Trim("! hel!lo!", " !") // 不会去掉中间的
	fmt.Printf("str=%v\n", str)

	//18、 \strings.TrimLeft() strings.TrimRight()将左边、右边的字符去掉

	//19、 判断字符串是否以指定的字符串开头
	b = strings.HasPrefix("ftp:192.168.1.1", "ftp")
	fmt.Printf("b=%v\n", b)
	
	//20、判断是否以指定字符串结尾
	// strings.HasSuffix()
}
