package Feistel

import "github.com/CookiesChen/Des-Golang/RoundFunction"

func interationT(preLR [64]int, key [48]int)(LR [64]int){
	preL, preR := [32]int{}, [32]int{}
	/* L(i)=R(i+1)*/
	for i:=0; i < 32; i++ {
		preL[i] = preLR[i]
		preR[i] = preLR[i+32]
		LR[i] = preLR[i+32]
	}
	/*  获取Ri*/
	f := RoundFunction.Function(preR, key)
	for i:=0; i < 32; i++ {
		LR[i+32] = preL[i] ^ f[i]
	}
	return LR
}