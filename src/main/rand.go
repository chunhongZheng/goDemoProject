package main

import (
	"fmt"
	"math/rand"
	"prlife/src/tool"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())                       //设置随机数种子，加上这行代码，可以保证每次随机都是随机的
	fmt.Println("My first lucky number is", rand.Intn(20)) //获取随机数，不加随机种子，每次遍历获取都是重复的一些随机数据
	fmt.Println("My senond lucky number is", rand.Intn(10))
	fmt.Println("My three lucky string is", tool.GetRandomString(10))

}
