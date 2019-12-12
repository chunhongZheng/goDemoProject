package main

import (
	"fmt"
	"math/big"
)

//base58和base64一样是一种二进制转可视字符串的算法，主要用来转换大整数值。区别是，转换出来的字符串，去除了几个看起来会产生歧义的字符，
// 如 0 (零), O (大写字母O), I (大写的字母i) and l (小写的字母L) ，和几个影响双击选择的字符，如/, +。
//
//结果字符集正好58个字符(包括9个数字，24个大写字母，25个小写字母)。
//base58的go实现   base58编码用于生成 比特币地址数据
func main() {

	//   6  %2== 3 ....0
	//   3%2==   1......1
	//    1%2===0 ......1       011      110

	//Base58Encode()

	//	var inputBtyes []byte=[0x00, 0xFF]
	//	inputBtyes=Base58Encode(inputBtyes)
	inputBtyes := []byte{0x01, 0xFF}
	result := Base58Encode(inputBtyes)
	//	println(result)
	println(1, 3, 4, 5)
	fmt.Printf("结果为: %s\n", result)
}

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) []byte {
	var result []byte                           //声明一个slice类型变量  切片
	x := big.NewInt(0).SetBytes(input)          //大数据类型
	base := big.NewInt(int64(len(b58Alphabet))) //相当于多少进制数
	zero := big.NewInt(0)
	mod := &big.Int{}
	for x.Cmp(zero) != 0 {
		//x是商,mod是余数    10/2=5 ....0     5/2=2 ....1    2/2=1....0    1/2=0....1          1010
		x.DivMod(x, base, mod) //对x取余数
		result = append(result, b58Alphabet[mod.Int64()])
		println(result)
	}
	//反转结果
	reverseByte(result)
	for _, b := range input {
		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]}, result...)

		} else {
			break
		}
	}
	return result
}

//  abcdef
//  1    fbcdea     i=0,j=5     data[0]=data[5]   data[5]=data[0]
//  2    fecdba     i=1,j=4
//  3  fedcba       i=2,j=3
//                  i=3,j=2 不满足条件，循环结束
//完成对调
func reverseByte(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

}
