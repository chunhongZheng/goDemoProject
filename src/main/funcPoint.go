package main

import "fmt"

func testFunc() {
	fmt.Printf("testFunc()被调用 \n")
}

//函数指针
func main() {
	var a = 7
	fmt.Printf("%d \n", a)
	testFunc()
	var test = testFunc //变量test的值其实就是函数testFunc的地址
	fmt.Printf("test的值为: %p\n", test)
	test()

}
