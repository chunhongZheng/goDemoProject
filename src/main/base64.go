package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

//base64编码测试类
//base64索引表：
// 数值
func main() {
	idno := "445"
	//比如445的ASCII码为 52 52 53     2进制码为：0011 0100   0011 0100  0011 0101  32+16+4          3*8      24位二进制
	//转换成64位    即  001101   000011 010000 110101    13（对应的索引表为N）   3（对应的索引表D）  16（对应的索引表Q）  53（对应的索引表1）   32+16+4+1=53       2*6 刚好是64     6位为一组
	//   的base64转变后就是 NDQ1
	idnoEncode := "NDQ1"
	fmt.Printf("%d\n", len(idno))
	fmt.Printf("%d\n", len(idnoEncode))
	src := "NDQ1"
	dst := baseDeEncode(src)
	println(dst) //解码后的数据为: 445

	//
	msg := []byte("i am the reader!")
	encodeMsg := baseStdEncode(msg)
	println(encodeMsg)

	dst2 := baseDeEncode(encodeMsg)
	println(dst2)
}

func baseStdEncode(srcBtye []byte) string {
	encoding := base64.StdEncoding.EncodeToString(srcBtye)
	return encoding
}

func baseDeEncode(src string) string {
	reader := strings.NewReader(src)
	decoder := base64.NewDecoder(base64.StdEncoding, reader)
	// 以流式解码
	buf := make([]byte, 2)
	// 保存解码后的数据
	dst := ""
	for {
		n, err := decoder.Read(buf)
		if n == 0 || err != nil {
			break
		}
		dst += string(buf[:n])
	}
	fmt.Println("解码后的数据为:", dst)
	return dst
}
