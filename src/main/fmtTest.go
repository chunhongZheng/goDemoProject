package main

import "fmt"

/***输出格式化**/
func main() {
	var x string = "caspar"
	//%T 类型
	fmt.Printf("%T,%v\n", x, x)
	//布尔类型
	var bool1 bool = true
	fmt.Printf("%T,%t\n", bool1, bool1)
	fmt.Printf("%T,%v\n", bool1, bool1)

	//整数
	fmt.Printf("%T,%d\n", 123, 123)
	fmt.Printf("%T,%5d\n", 123, 123)   //6代表长度
	fmt.Printf("%T,%06d\n", 123, 123)  //0代表填充的是0
	fmt.Printf("%T,%016d\n", 123, 123) //16代表长度，不够的前面填充0
	fmt.Printf("%d,%b\n", 10, 10)      //二进制
	fmt.Printf("%d,%o\n", 10, 10)      //八进制
	fmt.Printf("%d,%x\n", 10, 10)      //十六进制
	fmt.Printf("%d,%#0x\n", 123, 123)  //前面补0x, 小写 0x7b
	fmt.Printf("%d,%#0X\n", 123, 123)  //前面补0x, 小写 0X7B

}
