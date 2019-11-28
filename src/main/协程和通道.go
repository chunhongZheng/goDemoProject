package main

import (
	"fmt"
	"net/http"
)

//main方法 是一个主协程，当主协程执行完毕时，所有子协程都会立即中断，停止执行。
//
func main() {
	//定义链接切片
	links := []string{
		"http://www.baidu.com",
		"http://www.prlife.com.cn",
		"http://www.taobao.com",
		"http://www.163.com",
		"http://www.google.cn/",
		"http://www.google.cn/",
		"http://www.google.com.hk",
	}
	//fmt.Println(links)

	c := make(chan string)
	for _, link := range links {
		// fmt.Printf("链接地址:%v\n",link)
		//  CheckLink(link)   //当前程序不是并发执行，而是顺序执行，效率不高
		//添加go 关键字，每遍历一个链接，就新增一个协程。
		// 此时并没有想象中的输出，因为当循环执行结束后，主协程main方法就已经执行完毕，而子协程并没有执行完，故没有任何输出
		//  go CheckLink(link)

		go CheckLinkAndChannel(link, c)
	}
	//等待接收通道数据，当接收到第一个成功子协程返回的数据之后，主协程打印出通道信息数据，并结束执行。其他子协程也一并结束，故此时只会打印一个链接地址成功连接上的信息
	//	linkResult :=<-c     // 等待接收通道数据
	//close(c) // 关闭通道
	//   fmt.Println(linkResult)

	//等待所有子协程的执行返回，所有都执行结束之后才结束主协程的执行
	//打印结果如下：
	//待链接地址:http://www.baidu.com
	//待链接地址:http://www.google.com.hk
	//待链接地址:http://www.prlife.com.cn
	//待链接地址:http://www.163.com
	//待链接地址:http://www.taobao.com
	//待链接地址:http://www.google.cn/
	//待链接地址:http://www.google.cn/
	//链接地址http://www.163.com:已经成功连接上了
	//链接地址http://www.163.com:成功连接上了
	//链接地址http://www.baidu.com:已经成功连接上了
	//链接地址http://www.baidu.com:成功连接上了
	//链接地址http://www.google.cn/:已经成功连接上了
	//链接地址http://www.google.cn/:成功连接上了
	//链接地址http://www.google.cn/:已经成功连接上了
	//链接地址http://www.google.cn/:成功连接上了
	//链接地址http://www.prlife.com.cn:已经成功连接上了
	//链接地址http://www.prlife.com.cn:成功连接上了
	//链接地址http://www.taobao.com:已经成功连接上了
	//链接地址http://www.taobao.com:成功连接上了
	for _, _ = range links {
		linkResult := <-c
		fmt.Println(linkResult)
	}

}

//通道就是协程之间的通信
//channel可以传输基本类型的数据如int, string，同时也可以传输struct数据
func CheckLinkAndChannel(link string, c chan string) {
	//fmt.Printf("待链接地址:%v\n",link)
	_, err := http.Get(link)
	if err == nil {
		fmt.Println("链接地址" + link + ":已经成功连接上了")
		c <- "链接地址" + link + ":成功连接上了"
	} else {
		fmt.Println("链接地址" + link + "连接不上,当前网络不支持")
		c <- "连接不上,当前网络不支持"
	}
}

// 检查链接是否可以成功连接
func CheckLink(link string) {
	//resp, err:=http.Get(link)
	fmt.Printf("链接地址:%v\n", link)
	_, err := http.Get(link)
	//fmt.Printf("resp:%v\n,  err:%v\n",resp,err)
	if err == nil {
		fmt.Println("可以成功连接")
	} else {
		fmt.Println("连接不上,当前网络不支持")
	}
}
