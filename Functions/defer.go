package main

import "fmt"

// defer 在函数中执行完毕后，及时释放资源

// 当执行defer时，暂时不执行，会将defer 后面的语句压入到独立的栈
// 执行完毕后，再从defer 栈，按照 先入后出 的方式出栈
func sum(n1 int, n2 int) int {
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok1 n2=", n2)

	n1++ // n1 = 11
	n2++ // n2 = 21

	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}

// defer 使用细节
// 先入后出
// defer 将语句输入到栈，也会将相关值拷贝到栈中
/*

ok3 res= 32
ok1 n2= 20
ok1 n1= 10
res= 32


*/
