package main

import "fmt"

// golang 中基本类型转换
func main() {
	var i int32 = 100
	// 将 i = >  float32

	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Printf("i=%v, n1=%v, n2=%v, n3=%v ", i, n1, n2, n3)

	// 基本类型和string 转换

	var num1 int64 = 99
	var num2 float64 = 23.33
	var b bool = false
	var mychar byte = 'b'

	var str string // 空string
	//fmt.Sprintf("%参数",	表达式)

	str = fmt.Sprintf("%d", num1)                // 十进制表示
	fmt.Printf("str type %T str=%q\n", str, str) //%T 相应值的类型的Go语法表示 %q 单引号围绕的字符字面值，由Go语法安全地转义

	str = fmt.Sprintf("%f", num2) //%f     有小数点而无指数，例如 123.456
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b) //%t   true 或 false
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", mychar) //%c     相应Unicode码点所表示的字符 
	fmt.Printf("str type %T str=%q\n", str, str)

}
