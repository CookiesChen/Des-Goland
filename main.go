package main

import (
	"Des-Goland/Key"
	"fmt"
)
import "Des-Goland/Feistel"

var k = [64]int{1,1,1,1,1}

func main()  {
	keys := Key.GetKeys(k)
	c := Feistel.Cryption(k, keys)
	fmt.Printf("%v", c)
	fmt.Printf("%v", Feistel.Decrypt(c, keys))
}