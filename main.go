package main

import (
	_ "crypto/des"
	"fmt"
	"github.com/CookiesChen/Des-Golang/Feistel"
	"github.com/CookiesChen/Des-Golang/Key"
)

var k = [64]int{1,1,1,1}

func main()  {
	keys := Key.GetKeys(k)
	data := ""
	fmt.Scanln(&data)
	fmt.Printf("原始明文: %v\n",data)
	length := 8 - len(data) % 8
	for i := 0; i < length; i++ {
		data += string(length)
	}
	ciphertext := ""
	context := ""
	for i, times := 0, len(data)/8; i < times; i ++{
		block :=  byteToBit([]byte(data[8*i:8*(i+1)]))
		c := Feistel.Encrypt(block, keys)
		ciphertext += string(bitToByte(c))
		d := Feistel.Decrypt(c, keys)
		context += string(bitToByte(d))
	}
	last := int(context[len(context) - 1])
	context = context[:len(context) - last]
	fmt.Printf("密文: %v\n",ciphertext)
	fmt.Printf("解密后: %v\n",context)
}

func byteToBit(data []byte) [64]int {
	var dst [64]int
	for index, v := range data {
		for i := 0; i < 8; i++ {
			move := uint(7 - i)
			dst[index*8 + i] = int((v>>move)&1)
		}
	}
	return dst
}

func bitToByte(data [64]int) []byte {
	var dst []byte
	for i := 0; i < 8; i++ {
		var sum byte
		for j := 0; j < 8; j++ {
			move := uint(7 - j)
			sum += byte(data[i*8 + j] << move)
		}
		dst = append(dst, sum)
	}
	return dst
}