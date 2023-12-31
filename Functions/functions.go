
/*
1)值传递
2）引用传递
不管是值传递，还是引用传递，传递给函数的都是变量的副本，不同的是值传递的是值的拷贝，引用传递的是地址的拷贝。
一般来说地址的拷贝效率高，因为数据量小，因为数据量小，而值拷贝决定拷贝数据的大小，数据越大，效率越低

*/

/*
1）值类型: int, float, bool, string, 结构体
2）引用类型: 指针、切片、map、管道、interface
*/

/*
值传递： 变量直接存储值，内存通常在栈中分配
引用传递： 变量存储的是一个地址，这个地址对应的空间才是真正存储数据，内存通常在堆上分配。
当没有任何变量引用这个地址时，该地址对应的数据空间在数据空间上成为一个垃圾，由GC回收
*/




/*
函数的作用域

全局变量：首字母若为大写，那么作用域为整个程序
Name := "tom" 赋值语句不能在函数体外
*/
