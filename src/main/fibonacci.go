package main

import (
	"fmt"
	"math/big"
	"time"
)

//big.Int是golang提供的数据类型，用来处理比较大的整型数，当然它也可以处理诸如64bit，32bit的常用整数。

const LIM = 1000       //求第1000位的斐波那契数列
var fibs [LIM]*big.Int //使用数组保存计算出来的数列的指针
func main() {
	result := big.NewInt(0)
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("数列第 %d 位: %d\n", i+1, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("执行完成，所耗时间为: %s\n", delta)
}
func fibonacci(n int) (res *big.Int) {
	if n <= 1 {
		res = big.NewInt(1)
	} else {
		temp := new(big.Int)
		res = temp.Add(fibs[n-1], fibs[n-2])
	}
	fibs[n] = res
	return
}
