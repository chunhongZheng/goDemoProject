package main

import (
	"fmt"
	"math/big"
)

func printInt(args ...int) {

}

//对big包功能测试
func main() {
	println(1, 3, 4, 5)
	//	z1 := 123
	//	fmt.Print(z1)
	//func (z *Int) DivMod(x, y, m *Int) (*Int, *Int)
	//	inputBtyes := []byte{0x00, 0xFF}

	//  返回指针所对应的值
	//type Int struct {
	//	neg bool // sign
	//	abs nat  // absolute value of the integer
	//}
	// NewInt allocates and returns a new Int set to x.
	//func NewInt(x int64) *Int {
	//	return new(Int).SetInt64(x)
	//}

	//Go语言不支持运算符重载
	//big1 := new(big.Int).SetUint64(uint64(1000))
	//fmt.Println("big1 is: ", big1)
	//big2 := big1.Uint64()
	//fmt.Println("big2 is: ", big2)

	//c.Add(a, b)   计算总和 a + b 并将结果存储在 c 中，覆盖之前在 c 中保存的任何值。

	a := big.NewInt(10)
	b := big.NewInt(20)
	sum := big.NewInt(0)
	sum.Add(a, b)
	fmt.Printf("%d+%d的结果为: %d\n", a, b, sum)

	sub := big.NewInt(0)
	sub.Sub(a, b)
	fmt.Printf("%d-%d的结果为: %d\n", a, b, sub)

	mul := big.NewInt(0)
	mul.Mul(a, b)
	fmt.Printf("%d*%d的结果为: %d\n", a, b, mul)

	div := big.NewInt(0)
	div.Div(b, a)
	fmt.Printf("%d/%d的结果为: %d\n", b, a, div)
	mod := big.NewInt(0)
	mod1 := &big.Int{}
	//func (z *Int) DivMod(x, y, m *Int) (*Int, *Int)
	//DivMod 将 z 设置为商x div y和m至模x mod y并返回y(y，y)对(z，m)。如果y == 0，则发生除零运行时恐慌。
	b1 := big.NewInt(21)
	div.DivMod(b1, a, mod)
	div.DivMod(b1, a, mod1)
	fmt.Printf("%d/%d的商为: %d\n", b1, a, div)
	fmt.Printf("%d/%d的模为: %d\n", b1, a, mod)
	fmt.Printf("%d/%d的b1值为: %d\n", b1, a, b1)
}
