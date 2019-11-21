package main

import "fmt"

/***
指针 测试类

*/

//值传递, 只传递值，修改变量aa的值，其实相当于在修改一个新地址的变量值，并不会修改到原变量a的值
func change1(aa int) {
	aa += 2
	fmt.Printf("aa的值为%d, 地址为:%p \n", aa, &aa)
}

//地址传递
func change2(bb *int) {
	*bb += 2
	fmt.Printf("*bb的值为%d, 地址为:%p \n", *bb, bb)
}

func main() {
	//操作符& 为取地址运算符，如下方，&a的值即为变量a的地址0xc00000a0d8
	var a int = 72
	fmt.Printf("a 的值为 %d, a 的地址为: %p \n", a, &a)
	//操作符* (var b *int)的方式出现表示此变量为存储整型变量地址的变量，b的值是一个地址，  * 运算符有的叫解引用或解指针运算符
	//*b 表示取变量中所对应的值所对应地址存储的值。
	var b *int = &a
	fmt.Printf("b的值为: %p, *b所存储的值为 %d\n", b, *b)

	//值传递和地址传递
	change1(a)
	fmt.Printf("a 的值为 %d, a 的地址为: %p \n", a, &a)
	change2(b)
	fmt.Printf("b的值为: %p, *b所存储的值为 %d\n", b, *b)

}
