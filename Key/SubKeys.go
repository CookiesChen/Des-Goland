package Key

/* 获取子密钥 */
func GetKeys(k [64]int) (keys [16][48]int){
	/* PC-1置换获取C0,D0 */
	cd := pc1(k)
	for i:=0; i < 16; i++ {
		/* LS获取Ci,Di */
		cd = ls(cd, i+1)
		keys[i] = pc2(cd)
	}
	return keys
}