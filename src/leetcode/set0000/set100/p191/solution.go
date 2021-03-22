package p191

func hammingWeight(num uint32) int {
	var res int
	for num != 0 {
		res += int(num & 1)
		num >>= 1
	}
	return res
}
