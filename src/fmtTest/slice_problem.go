package fmtTest

import "fmt"

//函数名称需要大写
func TestSlicePTTroblem() {
	op := make([]*Person, 3)     // 注意：这是一种错误写法  make([]T, 3) 的含义是len=3，且 cap=3的数组，此时数组中已经存在3个nil的元素了
	op1 := make([]*Person, 0, 3) //正确写法，len=0,cap=3 不存在nil的元素
	fmt.Printf("新建 op类型为:%T op的值为%v,op的长度为:%d op的容量为%d\n", op, op, len(op), cap(op))
	fmt.Printf("新建 op1类型为:%T op1的值为%v,op1的长度为:%d op1的容量为%d\n", op1, op1, len(op1), cap(op1))
}
