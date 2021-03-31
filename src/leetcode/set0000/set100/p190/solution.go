package p190

func reverseBits(num uint32) uint32 {
	var res uint32
	var i int
	for num != 0 {
		res = ((num & 1) << (31 - i)) | res
		num >>= 1
		i++
	}
	return res
}
