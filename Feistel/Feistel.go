package Feistel

var keys =[16][48]int{}

func Encrypt(k [64]int, subKey [16][48]int)(output [64]int) {
	keys = subKey
	lr := ipdisplace(k)
	for i:=0; i < 16; i++ {
		lr = interationT(lr, keys[i])
	}
	turnLR := [64]int{}
	for i:=0; i < 32; i++ {
		turnLR[i] = lr[i+32]
		turnLR[i+32] = lr[i]
	}
	output = ipinverse(turnLR)
	return output
}

func Decrypt(k [64]int, subKey [16][48]int)(output [64]int){
	keys = subKey
	lr := ipdisplace(k)
	for i:=15; i >= 0; i-- {
		lr = interationT(lr, keys[i])
	}
	turnLR := [64]int{}
	for i:=0; i < 32; i++ {
		turnLR[i] = lr[i+32]
		turnLR[i+32] = lr[i]
	}
	output = ipinverse(turnLR)
	return output
}