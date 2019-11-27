package main

import (
	"fmt"
	"unsafe"
)

//定义结构体
type Struct_Person struct {
	Name    string
	Sex     string
	Age     int
	Address string
}

//修改结构体变量的某个属性值，将整个person变量传值进去，然后进行修改
func updateAddress(person Struct_Person) Struct_Person {
	person.Address = "天空一号"
	person.Name = "jimmy"
	//此处程序又开僻了一个新空间，修改的值也是新空间的值，原空间的值不变
	fmt.Printf("updateAddress()方法 person的地址为: %p \n", &person)
	fmt.Printf("updateAddress()方法新建的person的内存开销为:%d \n", unsafe.Sizeof(person))
	return person
}

//传递的是一个指针地址,返回修改过后的person
func updateAddressByPoint(person *Struct_Person) *Struct_Person {
	person.Address = "载人航天研究基地"
	person.Name = "john"
	//依然是指向同一个地址
	fmt.Printf("updateAddressByPoint()方法 person的地址为: %p \n", person)
	fmt.Printf("updateAddressByPoint()方法新建的person的内存开销为:%d \n", unsafe.Sizeof(person))
	return person
}

func main() {
	var person = Struct_Person{
		Name:    "caspar zheng",
		Sex:     "M",
		Age:     30,
		Address: "广州市天河区车陂",
	}
	fmt.Printf("person %v \n", person)
	fmt.Printf("main()方法 person的地址为: %p \n", &person)
	//var newPerson=updateAddress(person)
	//fmt.Printf("原来person %v,  地址为:%p \n",person,&person)
	//fmt.Printf("修改后的person %v  地址为:%p \n\n",newPerson,&newPerson)
	var newPerson *Struct_Person = updateAddressByPoint(&person) // 此处之所以地址不同，是因为重新定义了一个指针变量，所以此处地址不同
	fmt.Printf("原来person %v,  地址为:%p \n", person, &person)
	fmt.Printf("修改后的person %v  地址为:%p \n", newPerson, &newPerson)

}
