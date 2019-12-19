package main

import (
	"bytes"
	"compress/gzip"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"prlife/src/tool"
	"strings"
)

/* 先看一个栗子：
小明就读于小学二年级，会计算加法，但是不会计算除法。你是小明的怪蜀黍大强，你想出一道题给他做，让他虽然能理解题目意思但是做起来有难度：
强：“小明小明，过来，叔叔问你，1+1等于几？”
明：“叔叔的大学白念了吧，我幼儿园就会了，等于2。”
强：“那考你个难的，7+7等于几？”
明：“切，你当人家上课啃铅笔头，下课帮小红写作业都是白干的是吧。手指都不用数，等 于14呗。”
强：“行，有叔叔当年的风采，那叔叔再问你，几个7相加等于56？”
明：“……”，默默掏出草稿纸、铅笔、手指头、脚趾头，进行了10分钟的深度计算:2个7等于14，3个7等于21，4个7等于28……。“叔叔,我算出来了，是8个，对不对？”
强：“好小子，叔叔就不信考不倒你。几个7相加等于864192？” 你心中默念，以小明的计算能力，要算到这个数恐怕得一年半载的。
明：“叔叔好厉害呀，我算不出来。”
//明天去考考小红看她会不会算几个7等于56，不会算我就交她，嘿嘿。”
*/

/**
  通过一个随机key创建公钥和私钥
  随机key至少为36位
*/
func getEcdsaKey(randKey string) (*ecdsa.PrivateKey, ecdsa.PublicKey, error) {
	var err error

	var prk *ecdsa.PrivateKey
	var puk ecdsa.PublicKey
	var curve elliptic.Curve
	lenth := len(randKey)
	if lenth < 224/8 {
		err = errors.New("私钥长度太短，至少为36位！")
		return prk, puk, err
	}
	if lenth > 521/8+8 {
		curve = elliptic.P521()
	} else if lenth > 384/8+8 {
		curve = elliptic.P384()
	} else if lenth > 256/8+8 {
		curve = elliptic.P256()
	} else if lenth > 224/8+8 {
		curve = elliptic.P224()
	}
	//GenerateKey 生成公钥和私钥对。
	//func GenerateKey(c elliptic.Curve, rand io.Reader) (*PrivateKey, error)
	//ecdsa.GenerateKey(curve,rand.Reader)
	prk, err = ecdsa.GenerateKey(curve, strings.NewReader(randKey))
	if err != nil {
		return prk, puk, err
	}
	puk = prk.PublicKey //ECDSA 公钥
	return prk, puk, err
}

/**
  对text加密，text必须是一个hash值，例如md5、sha1等
  使用私钥prk
  使用随机熵增强加密安全，安全依赖于此熵，randsign
  返回加密结果，结果为数字证书r、s的序列化后拼接，然后用hex转换为string
*/
func sign(text []byte, randSign string, prk *ecdsa.PrivateKey) (string, error) {
	r, s, err := ecdsa.Sign(strings.NewReader(randSign), prk, text)
	if err != nil {
		return "", err
	}
	rt, err := r.MarshalText()
	if err != nil {
		return "", err
	}
	st, err := s.MarshalText()
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	_, err = w.Write([]byte(string(rt) + "+" + string(st)))
	if err != nil {
		return "", err
	}
	w.Flush()
	return hex.EncodeToString(b.Bytes()), nil
}

/**
  证书分解
  通过hex解码，分割成数字证书r，s
*/
func getSign(signature string) (rint, sint big.Int, err error) {
	byterun, err := hex.DecodeString(signature)
	if err != nil {
		err = errors.New("decrypt error, " + err.Error())
		return
	}
	r, err := gzip.NewReader(bytes.NewBuffer(byterun))
	if err != nil {
		err = errors.New("decode error," + err.Error())
		return
	}
	defer r.Close()
	buf := make([]byte, 1024)
	count, err := r.Read(buf)
	if err != nil {
		fmt.Println("decode = ", err)
		err = errors.New("decode read error," + err.Error())
		return
	}
	rs := strings.Split(string(buf[:count]), "+")
	if len(rs) != 2 {
		err = errors.New("decode fail")
		return
	}
	err = rint.UnmarshalText([]byte(rs[0]))
	if err != nil {
		err = errors.New("decrypt rint fail, " + err.Error())
		return
	}
	err = sint.UnmarshalText([]byte(rs[1]))
	if err != nil {
		err = errors.New("decrypt sint fail, " + err.Error())
		return
	}
	return

}

/**
  校验文本内容是否与签名一致
  使用公钥校验签名和文本内容
*/
func verify(text []byte, signature string, key ecdsa.PublicKey) (bool, error) {

	rint, sint, err := getSign(signature)
	if err != nil {
		return false, err
	}
	result := ecdsa.Verify(&key, text, &rint, &sint)
	return result, nil

}

/**
  hash加密
  使用md5加密
*/
func hashtext(text, salt string) []byte {

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(text))
	result := Md5Inst.Sum([]byte(salt))

	return result
}

func main() {
	//随机熵，用于加密安全
	randSign := "20180619zafes20180619zafes20180619zafessss" //至少36位
	//随机key，用于创建公钥和私钥
	randKey := tool.GetRandomString(40)
	//randKey := "fb0f7279c18d4394594fc9714797c9680335a320"
	//创建公钥和私钥
	prk, puk, err := getEcdsaKey(randKey)
	if err != nil {
		fmt.Println(err)
	}
	publicKey := append(puk.X.Bytes(), puk.Y.Bytes()...)
	fmt.Printf("私钥为：%x 长度为:%d\n公钥为：%x,公钥长度为%d\n", prk.D.Bytes(), len(prk.D.Bytes()), publicKey, len(publicKey))
	//hash加密使用md5用到的salt
	salt := "131ilzaw"
	//待加密的明文
	text := "hlloaefaefaefaefaefaefaefhelloaefaefaefaefaefaefaefhelloaefaefaefaefaefaefaef"
	//hash取值
	htext := hashtext(text, salt)
	//hash值编码输出
	fmt.Println(hex.EncodeToString(htext))
	//hash值进行签名
	result, err := sign(htext, randSign, prk)
	if err != nil {
		fmt.Println(err)
	}
	//签名输出
	fmt.Println(result)

}
