package p5626

func minPartitions(str string) int {
	var res int
	var prev int
	for i := 0; i < len(str); i++ {
		cur := int(str[i] - '0')
		if cur <= prev {
			continue
		}
		res += cur - prev
		prev = cur
	}
	return res
}
