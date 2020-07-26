package p1527

func minFlips(target string) int {
	n := len(target)
	var res int
	for i := 0; i < n; i++ {
		if int(target[i]-'0') != res&1 {
			res++
		}
	}
	return res
}
