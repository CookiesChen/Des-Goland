package RoundFunction

func eExtend(R [32]int)(E [48]int)  {
	table := [48]int {
		32, 1,  2,  3,  4,  5,
		4,  5,  6,  7,  8,  9,
		8,  9,  10, 11, 12, 13,
		12, 13, 14, 15, 16, 17,
		16, 17, 18, 19, 20, 21,
		20, 21, 22, 23, 24, 25,
		24, 25, 26, 27, 28, 29,
		28, 29, 30, 31, 32, 1}

	for i:= 0; i< 8; i++ {
		for j:= 0; j < 6; j++ {
			index := i*6 + j
			E[index] = R[table[index] - 1]
		}
	}
	return E
}