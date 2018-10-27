package main

import (
	"Des-Goland/Key"
	_ "crypto/des"
	"fmt"
)
import "Des-Goland/Feistel"

var k = [64]int{1,1,1,1}

func main()  {
	keys := Key.GetKeys(k)
	c := Feistel.Encrypt(k, keys)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v", Feistel.Decrypt(c, keys))
}