package p2085

func minimumBuckets(street string) int {
	n := len(street)
	var res int
	prev := -2
	for i := 0; i < n; i++ {
		if street[i] == 'H' {
			// must place at next position
			if prev == i-1 {
				continue
			}
			if i+1 < n && street[i+1] == '.' {
				// can place after
				res++
				prev = i + 1
				continue
			}
			if i == 0 || street[i-1] == 'H' {
				//can't place before
				return -1
			}
			res++
			prev = i - 1
		}
	}

	return res
}
