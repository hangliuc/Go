package main

import (
	"fmt"
	"time"
)

func main() {
	// 1、获取当前时间
	//类型为 type=time.Time
	now := time.Now() //获得时间方法
	fmt.Printf("now=%v now\ntype=%T\n", now, now)

	//2、方法获取详细
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//3、格式化日期时间
	//方式一
	fmt.Printf("当前年月日: %d-%d-%d %d:%d:%d",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	//方式二
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("15")) //只取出 “时”

	// 4、时间的常量
	// 在程序中获得指定时间的单位
	//const (
	//	Nanosecond  Duration = 1
	//	Microsecond          = 1000 * Nanosecond
	//	Millisecond          = 1000 * Microsecond
	//	Second               = 1000 * Millisecond
	//	Minute               = 60 * Second
	//	Hour                 = 60 * Minute
	//)

	// 5、sleep 休眠
	//每隔一秒钟打印一个数字

	i := 0
	for {
		i++
		fmt.Println(i)
		//休眠
		time.Sleep(time.Millisecond * 100) // 休眠0.1秒
		// 休眠 1 s time.Sleep(time.Second)
		if i == 10 {
			break
		}
	}

	//6、获得当前unix 时间戳和unixnaon时间戳
	//可以获得随机数字

	fmt.Printf("unix 时间戳=%v unixnano 时间戳=%v", now.Unix(), now.UnixNano())
	
}


// 查询函数执行的时间
//package main
//
//import (
//	"fmt"
//	"strconv"
//	"time"
//)
//
//func test03() {
//	str := ""
//	for i := 0; i < 100000; i++ {
//		str += "hello" + strconv.Itoa(i)
//	}
//}
//
//func main() {
//	start := time.Now().Unix()
//	test03()
//	end := time.Now().Unix()
//	fmt.Printf("执行test03耗费的时间%v秒\n", end-start)
//}
