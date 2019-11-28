package main

import "fmt"

//range测试类
//range关键字是Go语言中一个非常有用的迭代array，slice，map, string, channel中元素的内置关键字。

//  for int i=0;i<5;i++{ }
func main() {
	nums := []int{2, 3, 4}
	sum := 0
	//用range子句替换掉for子句。range子句包含一个或两个迭代变量（用于与迭代出的值绑定）、特殊标记:=或=、关键字range以及range表达式。
	// 其中，range表达式的结果值的类型应该是能够被迭代的，包括：字符串类型、数组类型、数组的指针类型、切片类型、字典类型和通道类型
	for _, num := range nums {
		//slice []E   第一个参数index int类型 第二个参数value E[i]
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums {
		if num == 4 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana", "c": "orange", "d": "tomato"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}
	//rangeString()
	//rangeStringMore()
	//	funParameterTest(1,"func  test")
	funParameterTest2("a", "b")
}

func modifySlice() {
	v := []int{1, 2, 3, 4}
	for i := range v {
		v = append(v, i)
		fmt.Printf("Modify Slice: value:%v\n", v)
	}
	//range在执行之初就构建好了v的长度，即是4次，无论如何后面v的长度如何发生变化，都只会执行4次。所以该示例并不会导致死循环执行下去
	//打印结果
	//Modify Slice: value:[1 2 3 4 0]
	//Modify Slice: value:[1 2 3 4 0 1]
	//Modify Slice: value:[1 2 3 4 0 1 2]
	//Modify Slice: value:[1 2 3 4 0 1 2 3]
}

//map 的迭代顺序不固定，随机
func modifyMap() {
	data := map[string]string{"a": "A", "b": "B", "c": "C"}
	for k, v := range data {
		data[v] = k
		fmt.Println("modify Mapping", data)
	}
}

//string
//迭代的是Unicode不是字节，第一个返回值是UTF-8编码第一个字节的索引，所以索引值可能并不是连续的
//第二个返回值的类型为rune，不是string类型的，如果要使用该值需要格式化
func rangeString() {
	datas := "aAbB"

	for k, d := range datas {
		fmt.Printf("k的值为:%v,d的值(Unicode码)为:%d, d的字符值为:%c\n", k, d, d)
		//	fmt.Printf("k_addr:%p, k_value:%v\nd_addr:%p, d_value:%v\n----\n", &k, k, &d, d)
	}
}

//这段代码使用range迭代字符串"中\x80文"，字符串中\x80是一个无效的的Unicode字符，所以range在迭代时会使用U+FFFD将其替换。
// 另外UTF-8使用变长方式编码，第一个汉字中占用了3个字节，所以遍历第二个字符的时候，其索引已经是3了，但是其只占一个字节，所以第三个字符文的索引4.
func rangeStringMore() {
	for pos, char := range "中\x80文" { // \x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}
}

func funParameterTest(a int, b string) {
	fmt.Printf("%v, %v", a, b)
}

func funParameterTest2(a, b string) {
	fmt.Printf("%v, %v", a, b)
}
