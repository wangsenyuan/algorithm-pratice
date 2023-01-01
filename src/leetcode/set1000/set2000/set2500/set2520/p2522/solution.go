package p2522

func minimumPartition(s string, k int) int {
	var res int

	for i := 0; i < len(s); {
		res++
		var num int

		j := i
		for i < len(s) {
			num = num*10 + int(s[i]-'0')
			if num > k {
				break
			}
			i++
		}
		if i == j {
			return -1
		}
	}
	return res
}
