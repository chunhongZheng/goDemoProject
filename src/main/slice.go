package main

import (
	"fmt"
)

//内置引用类型slice
func main() {
	//var arr1[3] int
	//fmt.Printf("arr1=%v \n", arr1)
	//var arr = [4]int{1, 2, 3, 4}
	//fmt.Printf("arr=%v \n", arr)

	// var arr1=[3]int{1,2,3}
	//fmt.Printf("数组arr1=%v \n", arr1)
	////切片的声明
	//var slice1 =[] int{1,2,3,4,5,6,7}
	//fmt.Printf("slice1=%v \n", slice1)
	//  func make([]T, len, cap) []T
	//其中T代表要创建的切片的元素类型。 make函数采用类型，长度和可选容量。 调用时，make会分配一个数组并返回一个引用该数组的切片。
	//make会分配一个数组并返回一个引用该数组的切片。
	// slice2 :=  make([]int,5,7)
	//
	//fmt.Printf("slice2=%v,长度为:%d 容量为:%d \n", slice2,len(slice2),cap(slice2))
	//
	//d := []byte{'r', 'o', 'a', 'd'}
	//e := d[2:]
	//fmt.Printf("d=%v,长度为:%d 容量为:%d \n", d,len(d),cap(d))
	//fmt.Printf("e=%v,长度为:%d 容量为:%d \n", e,len(e),cap(e))
	rangeInt := make([]int, 5, 5)
	for i := range rangeInt {
		rangeInt[i] = i
	}
	//	s := make([]string, 5)
	s := []string{"r", "o", "a", "d"}
	s1 := s[2:] //切片不会复制切片的数据。 它创建一个指向原始数组的新切片值
	//	fmt.Printf("切片 s1的值为%v,s1的长度为:%d ,s1的容量为:%d\n",s1,len(s1),cap(s1))
	s1[0] = "m" //
	//	fmt.Printf("修改过后切片 s1的值为%v,s1的长度为:%d ,s1的容量为:%d\n",s1,len(s1),cap(s1))
	//修改重新切片的元素（而不是切片本身）会修改原始切片的元素
	//	fmt.Printf("修改过后切片 s的值为%v,s的长度为:%d ,s的容量为:%d\n",s,len(s),cap(s))
	//	s = s[2:4]
	//要增加切片的容量，必须创建一个新的更大的切片并将原始切片的内容复制到切片中。
	//长度是切片引用的元素数。 容量是底层数组中元素的数量（从切片指针引用的元素开始）
	//	fmt.Printf("原先 s的值为%v,s的长度为:%d ,s的容量为:%d\n",s,len(s),cap(s))
	//Slice增长（复制和追加功能）
	//t := make([]string, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
	//fmt.Printf("新建 t的值为%v,t的长度为:%d t的容量为%d\n",t,len(t),cap(t))
	//range的使用非常简单，对于遍历array，*array，string它返回两个值分别是数据的索引和值，遍历map时返回的两个值分别是key和value，
	// 遍历channel时，则只有一个返回数据
	//for i := range s {
	//	t[i] = s[i]
	//}
	//fmt.Printf("复制后t的值为%v,t的容量为%d\n",t,cap(t))
	//s = t
	//fmt.Printf("复制后s的值为%v,s的容量为:%d\n",s,cap(s))

	//使用copy可以简化上面的代码
	//func copy(dst, src []T) int    从源数据src 复制到目标数据dst
	t2 := make([]string, len(s), (cap(s)+1)*2)
	copy(t2, s)
	//	fmt.Printf("原先 s的值为%v,s的长度为:%d ,s的容量为:%d\n",s,len(s),cap(s))
	s = t2
	//	fmt.Printf("新建 t2的值为%v,t2的长度为:%d t2的容量为%d\n",t2,len(t2),cap(t2))
	//	fmt.Printf("copy后 s的值为%v,s的长度为:%d ,s的容量为:%d\n",s,len(s),cap(s))
	//表达式b [1：4]创建包括b的元素1到3的切片（得到的切片的索引将是0到2）
	//切片表达式的开始和结束索引是可选的; 它们分别默认为零和切片长度：
	//b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	// b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
	//切片的零值为nil。 对于nil，len和cap函数都将返回0。
	nilTest := []int{}
	fmt.Printf("新建 nilTest类型为:%T nilTest的值为%v,nilTest的长度为:%d nilTest的容量为%d\n", nilTest, nilTest, len(nilTest), cap(nilTest))
	//
	p := []byte{2, 3, 5}
	p = append(p, 1)
	p = append(p, 2, 3, 4, 5, 6)
	fmt.Printf("新建 p类型为:%T p的值为%v,p的长度为:%d p的容量为%d\n", p, p, len(p), cap(p))
	//p = AppendByte(p, 7, 11, 13)
	//fmt.Printf("新建 p类型为:%T p的值为%v,p的长度为:%d p的容量为%d\n",p,p,len(p),cap(p))
	//p = AppendByte(p,8,6,12)
	//fmt.Printf("新建 p类型为:%T p的值为%v,p的长度为:%d p的容量为%d\n",p,p,len(p),cap(p))
	//// p == []byte{2, 3, 5, 7, 11, 13}
	////GO 内置的附加功能  append函数将元素x附加到切片s的末尾，如果需要更大的容量，则增加切片。
	////func append(s []T, x ...T)
	// p=append(p,1,1,1)
	//fmt.Printf("新建 p类型为:%T p的值为%v,p的长度为:%d p的容量为%d\n",p,p,len(p),cap(p))
	//p=append(p,2,2,2,3,3)
	//fmt.Printf("新建 p类型为:%T p的值为%v,p的长度为:%d p的容量为%d\n",p,p,len(p),cap(p))
	//
	//var helloString ="helloWorld"
	//fmt.Printf("%v \n",helloString)

	//要将一个切片附加到另一个切片，请使用...将第二个参数展开为参数列表。
	slice_a := []string{"John", "Paul"}
	slice_b := []string{"George", "Ringo", "Pete"}
	slice_a = append(slice_a, slice_b...)
	fmt.Printf("新建 slice_a类型为:%T slice_a的值为%v,slice_a的长度为:%d slice_a的容量为%d\n", slice_a, slice_a, len(slice_a), cap(slice_a))
	//a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
	// a == []string{"John", "Paul", "George", "Ringo", "Pete"}

}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	//必要时更新切片，并将更新后的切片返回给原切片，相当于扩容
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
