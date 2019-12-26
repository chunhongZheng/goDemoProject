package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"
	"golang.org/x/crypto/ripemd160"
	"math/big"
)

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
const VERSION = byte(0x00)
const CHECKSUM_LENGTH = 4

type BitcoinKeys struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
}


func GetBitcoinKeys() *BitcoinKeys {
	b := &BitcoinKeys{nil, nil}
	b.newKeyPair()
	return b
}

func (b *BitcoinKeys) newKeyPair() {
	curve := elliptic.P256()
	var err error
	b.PrivateKey, err = ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	b.PublicKey = append(b.PrivateKey.PublicKey.X.Bytes(), b.PrivateKey.PublicKey.Y.Bytes()...)
}

func GeneratePublicKeyHash(publicKey []byte) []byte {
	sha256PubKey := sha256.Sum256(publicKey)
	r := ripemd160.New()
	r.Write(sha256PubKey[:])
	ripPubKey := r.Sum(nil)
	return ripPubKey
}
//通过地址获得公钥
func GetPublicKeyHashFromAddress(address string) []byte {
	addressBytes := []byte(address)
	fullHash := Base58Decode(addressBytes)
	publicKeyHash := fullHash[1 : len(fullHash)-CHECKSUM_LENGTH]
	return publicKeyHash
}
func CheckSumHash(versionPublickeyHash []byte) []byte {
	versionPublickeyHashSha1 := sha256.Sum256(versionPublickeyHash)
	versionPublickeyHashSha2 := sha256.Sum256(versionPublickeyHashSha1[:])
	tailHash := versionPublickeyHashSha2[:CHECKSUM_LENGTH]
	return tailHash
}

//检测比特币地址是否有效
func IsVaildBitcoinAddress(address string) bool {
	adddressByte := []byte(address)
	fullHash := Base58Decode(adddressByte)
	if len(fullHash) != 25 {
		return false
	}
	prefixHash := fullHash[:len(fullHash)-CHECKSUM_LENGTH]
	tailHash := fullHash[len(fullHash)-CHECKSUM_LENGTH:]
	tailHash2 := CheckSumHash(prefixHash)
	if bytes.Compare(tailHash, tailHash2[:]) == 0 {
		return true
	} else {
		return false
	}
}

//获取地址

func (b *BitcoinKeys) GetAddress() []byte{
	//1.ripemd160(sha256(publickey))
	ripPubKey := GeneratePublicKeyHash(b.PublicKey)
	//2.最前面添加一个字节的版本信息获得 versionPublickeyHash
	versionPublickeyHash := append([]byte{VERSION}, ripPubKey[:]...)
	//3.sha256(sha256(versionPublickeyHash))  取最后四个字节的值
	tailHash := CheckSumHash(versionPublickeyHash)
	//4.拼接最终hash versionPublickeyHash + checksumHash
	finalHash := append(versionPublickeyHash, tailHash...)
	//进行base58加密
	address := Base58Encode(finalHash)
	return address
}

////加密
func Base58Encode(input []byte) []byte {
	var result []byte                           //声明一个slice类型变量  切片
	x := big.NewInt(0).SetBytes(input)          //大数据类型
	base := big.NewInt(int64(len(b58Alphabet))) //相当于多少进制数
	zero := big.NewInt(0)
	mod := &big.Int{} //大整数的指针
	for x.Cmp(zero) != 0 {
		//x是商,mod是余数    10/2=5 ....0     5/2=2 ....1    2/2=1....0    1/2=0....1          1010
		x.DivMod(x, base, mod) //对x取余数
		result = append(result, b58Alphabet[mod.Int64()])
		//println(result)
	}
	//反转结果
	reverseByte(result)
	//把这个字节前置的0 字节置换成第一个字符
	for _, b := range input {
		if b == 0x00 {

			result = append([]byte{b58Alphabet[0]}, result...)

		} else {
			break
		}
	}
	return result
}

//解密
func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0 //统计前面为0的个数，方便截取解码
	for _, b := range input {
		if b != b58Alphabet[0] {
			break
		}
		zeroBytes++
	}
	payload := input[zeroBytes:]
	for _, b := range payload {
		charIndex := bytes.IndexByte(b58Alphabet, b) //找出b所在的索引位置，即是余数
		result.Mul(result, big.NewInt(int64(len(b58Alphabet))))
		result.Add(result, big.NewInt(int64(charIndex)))
	}
	decoded := result.Bytes()
	decoded = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decoded...)
	return decoded
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

func main(){

}




