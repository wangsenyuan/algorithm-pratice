package p2491

func dividePlayers(skill []int) int64 {
	n := len(skill)
	var sum int
	rec := make(map[int]int)
	for _, sk := range skill {
		sum += sk
		rec[sk]++
	}
	h := n / 2
	if sum%h != 0 {
		return -1
	}
	team := sum / h
	var res int64

	for i := 0; i < n; i++ {
		x := skill[i]
		if rec[x] == 0 {
			continue
		}
		y := team - x
		if x != y {
			if rec[y] == 0 {
				return -1
			}
			rec[x]--
			rec[y]--
			res += int64(x) * int64(y)
		} else {
			if rec[x] < 2 {
				return -1
			}
			rec[x] -= 2
			res += int64(x) * int64(x)
		}
	}

	return res
}
