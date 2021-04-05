package p1814

func maxHappyGroups(batchSize int, groups []int) int {
	freq := make([]int, batchSize)
	for i := 0; i < len(groups); i++ {
		freq[groups[i]%batchSize]++
	}
	w := make([]int, batchSize+1)
	w[1] = 1
	for i := 2; i < batchSize; i++ {
		w[i] = w[i-1] * (freq[i-1] + 1)
	}
	sum := 1
	for i := 1; i < batchSize; i++ {
		sum *= (freq[i] + 1)
	}
	f := make([]int, sum)

	fq := make([]int, batchSize)

	for mask := 1; mask < sum; mask++ {
		var last int
		for i := 1; i < batchSize; i++ {
			fq[i] = (mask / w[i]) % (freq[i] + 1)
			last = (last + (freq[i]-fq[i])*i) % batchSize
		}
		for b := 1; b < batchSize; b++ {
			if fq[b] > 0 {
				tmp := f[mask-w[b]]
				if last == 0 {
					tmp++
				}
				f[mask] = max(f[mask], tmp)
			}
		}
	}
	return f[sum-1] + freq[0]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
