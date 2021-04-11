package lcp33

func storeWater(bucket []int, vat []int) int {
	var n int
	for i := 0; i < len(bucket); i++ {
		if vat[i] > 0 {
			bucket[n] = bucket[i]
			vat[n] = vat[i]
			n++
		}
	}
	if n == 0 {
		return 0
	}

	find := func(y int) int {
		// 检查操作2，使用最多y次的情况下, 所需要的总数量
		var res int
		for i := 0; i < n; i++ {
			b := bucket[i]
			v := vat[i]
			// (b + x) * y >= v
			// (b + x) >= v / y
			bb := v / y
			if bb*y < v {
				bb++
			}
			res += max(0, bb-b)
		}
		return res + y
	}

	best := 1 << 30

	mx := 1
	var fill int
	for i := 0; i < n; i++ {
		if bucket[i] == 0 {
			fill++
			bucket[i]++
		}
		mx = max(mx, vat[i]/bucket[i]+1)
	}

	for y := 1; y <= mx; y++ {
		tmp := find(y)
		if tmp < best {
			best = tmp
		}
	}
	return best + fill
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
