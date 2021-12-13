package p2103

func countPoints(rings string) int {
	n := len(rings) / 2
	flags := make([]int, 10)

	for i := 0; i < n; i++ {
		c := rings[2*i]
		j := int(rings[2*i+1] - '0')
		var k int
		if c == 'R' {
			k = 0
		} else if c == 'G' {
			k = 1
		} else {
			k = 2
		}
		flags[j] |= 1 << k
	}
	var res int
	for i := 0; i < 10; i++ {
		if flags[i] == 7 {
			res++
		}
	}
	return res
}
