package fmtTest

import (
	"fmt"
)

//var  x int=1
var s = "scopeTestString"

//定义结构体
type Person struct {
	Name    string
	Sex     string
	Age     int
	Address string
}

func PrintGeneralFmt() {

	var a = 8
	//%v	the value in a default format
	//	when printing structs, the plus flag (%+v) adds field names
	//fmt.Printf("%v \n",a)  //默认的格式   8
	//fmt.Printf("s: %v \n",s)  //string的打印格式  scopeTestString
	var person = Person{
		Name:    " fmt caspar zheng",
		Sex:     "M",
		Age:     30,
		Address: "广州市天河区车陂",
	}
	//结构体对象person数据为：{Name: fmt caspar zheng Sex:M Age:30 Address:广州市天河区车陂}
	//fmt.Printf("结构体对象person数据为：%+v \n",person)  //{Name: fmt caspar zheng Sex:M Age:30 Address:广州市天河区车陂}

	//%#v	a Go-syntax representation of the value
	//fmt.Printf("%#v \n",a)  //默认的格式
	//fmt.Printf("s: %#v \n",s)  //string的打印格式
	//结构体对象person数据为：fmtTest.Person{Name:" fmt caspar zheng", Sex:"M", Age:30, Address:"广州市天河区车陂"}
	//fmt.Printf("结构体对象person数据为：%#+v \n",person)

	//%T	a Go-syntax representation of the type of the value
	//打印值的类型
	fmt.Printf("a的值为:%#v a类型为: %T\n", a, a)
	fmt.Printf("s的值为:%#v s类型为: %T\n", s, s)
	fmt.Printf("结构体对象person值为: %#v  类型为: %T\n", person, person)
	//%%   不消耗任何值  打印%
	fmt.Printf(" %% \n")
}

func PrintBoolFmt() {
	var bool = false
	//%t	the word true or false
	fmt.Printf("bool变量的值为:%t", bool)
}

//整数值：
//
//%b	二进制表示
//%c	相应Unicode码点所表示的字符
//%d	十进制表示
//%o	八进制表示
//%q	单引号围绕的字符字面值，由Go语法安全地转义
//%x	十六进制表示，字母形式为小写 a-f
//%X	十六进制表示，字母形式为大写 A-F
//%U	Unicode格式：U+1234，等同于 "U+%04X"
func PrintIntegerFmt() {
	var x = 8
	// %b 二进制   %d 十进制
	fmt.Printf("%d的二进制代码为:%b \n", x, x)
	//%c  相应Unicode码点所表示的字符
	fmt.Printf("%c\n", 33)
	fmt.Printf("%c \n", 0x4E2D) //中  中的Unicode码为0x4E2D
	var unicode = "\u4e2d\u56fd\u68a6"
	fmt.Printf("%c\n", unicode)
	//fmt.Printf("%d相对应的Unicode码相对应的字符为为:%c \n",y,y)
	//%x 提供十六进制编码
	fmt.Printf("%x\n", 17)
	fmt.Printf("%x\n", 456)
	// %o 八进制   %d 十进制
	fmt.Printf("%d的八进制代码为:%o \n", x, x)
}
