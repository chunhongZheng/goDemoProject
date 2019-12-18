package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
	"fmt"
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
	puk = prk.PublicKey
	return prk, puk, err
}

func main() {
	//随机熵，用于加密安全
	//randSign := "20180619zafes20180619zafes20180619zafessss"//至少36位
	//随机key，用于创建公钥和私钥

	randKey := "fb0f7279c18d4394594fc9714797c9680335a320"
	//创建公钥和私钥
	prk, puk, err := getEcdsaKey(randKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("私钥为：%d,公钥为：%d", prk, puk)
}
