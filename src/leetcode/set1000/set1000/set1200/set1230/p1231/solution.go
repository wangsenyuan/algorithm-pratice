package p1231

func maximizeSweetness(sweetness []int, K int) int {
	// n := len(sweetness)
	var sum int
	for _, sweet := range sweetness {
		sum += sweet
	}

	check := func(x int) bool {
		if x*K < 0 {
			return false
		}
		if x*K > sum {
			return false
		}
		var k int
		var y int
		for i := 0; i < len(sweetness); i++ {
			y += sweetness[i]
			if y >= x {
				k++
				y = 0
			}
		}
		if y >= x {
			k++
		}
		return k > K
	}

	left, right := 0, sum+1

	for left < right {
		mid := (left + right) / 2
		if !check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right - 1
}
