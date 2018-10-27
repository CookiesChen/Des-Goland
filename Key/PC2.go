package Key

/* PC-2压缩置换 */
func pc2(k [56]int) (output [48]int){

	table := [48]int {
		14, 17, 11, 24, 1,  5,
		3,  28, 15, 6,  21, 10,
		23, 19, 12, 4,  26, 8,
		16, 7,  27, 20, 13, 2,
		41, 52, 31, 37, 47, 55,
		30, 40, 51, 45, 33, 48,
		44, 49, 39, 56, 34, 53,
		46, 42, 50, 36, 29, 32}

	temp := [56]int{}
	for i:= 0; i< 8; i++ {
		for j:= 0; j < 6; j++ {
			index := i*6 + j
			temp[index] = k[table[index] - 1]
		}
	}
	count := 0
	for i:= 0; i< 56; i++ {
		j := i + 1
		if !(j == 9 || j == 18 || j == 22 || j == 25 ||
			j == 35 || j == 38 || j == 43 || j == 54) {
			output[count] = temp[i]
			count++
		}
	}
	return output
}