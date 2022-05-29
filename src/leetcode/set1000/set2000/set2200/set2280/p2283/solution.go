package p2283

func digitCount(num string) bool {
	cnt := make([]int, 10)
	for _, i := range num {
		cnt[int(i-'0')]++
	}

	for i := 0; i < len(num); i++ {
		if cnt[i] != int(num[i] - '0') {
			return false
		}
	}
	return true
}
