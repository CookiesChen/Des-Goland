package Key

/* LSi获取Ci,Di  */
func ls(cd [56]int, index int) (output [56]int) {
	len := 1
	if index == 1 || index == 2 ||
		index == 6 || index == 16 {
			len = 2
	}
	for i:=0; i < 56; i++ {
		output[i] = cd[(i-len+56)%56]
	}
	return output
}