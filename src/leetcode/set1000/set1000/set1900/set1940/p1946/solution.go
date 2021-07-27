package p1946

func maximumNumber(num string, change []int) string {

	buf := []byte(num)

	for i := 0; i < len(buf); i++ {
		x := int(buf[i] - '0')
		if change[x] > x {
			// better to change it
			for j := i; j < len(buf); j++ {
				y := int(buf[j] - '0')
				if change[y] < y {
					break
				}
				buf[j] = byte(change[y] + '0')
			}
			break
		}
	}

	return string(buf)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
