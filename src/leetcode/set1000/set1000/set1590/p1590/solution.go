package p1590

func minSubarray(nums []int, p int) int {
	// (sum - sub) % p == 0 => sum % p == sub % p
	n := len(nums)
	var sum int64
	for i := 0; i < n; i++ {
		sum += int64(nums[i])
	}
	P := int64(p)
	R := sum % P
	if R == 0 {
		return 0
	}
	r := int(R)
	best := -1
	// find some [i, j] => sum(i, j) % p == r
	pos := make(map[int]int)
	pos[0] = -1
	sum = 0
	for i := 0; i < n; i++ {
		sum += int64(nums[i])
		cur := int(sum % P)
		tmp := cur - r
		if tmp < 0 {
			tmp += p
		}
		if j, found := pos[tmp]; found {
			if best < 0 || i-j < best {
				best = i - j
			}
		}
		pos[cur] = i
	}

	if best == n {
		return -1
	}
	return best
}
