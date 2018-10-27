package RoundFunction

func Function(R [32]int, key [48]int)(output [32]int){
	E := eExtend(R)
	output = pReplace(sBox(xor(E, key)))
	return output
}

func xor(E [48]int, key [48]int)(output [8][6]int)  {
	for i:=0; i < 8; i++ {
		for j:=0; j < 6; j++{
			index := i*6 + j
			output[i][j] = E[index] ^ key[index]
		}
	}
	return output
}