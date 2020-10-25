package p5546

func slowestKey(releaseTimes []int, keysPressed string) byte {
	time := make([]int, 26)

	for i := 0; i < len(releaseTimes); i++ {
		take := releaseTimes[i]
		if i > 0 {
			take -= releaseTimes[i-1]
		}
		x := int(keysPressed[i] - 'a')
		time[x] = max(time[x], take)
	}
	var best int
	for i := 1; i < 26; i++ {
		if time[i] >= time[best] {
			best = i
		}
	}
	return byte('a' + best)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
