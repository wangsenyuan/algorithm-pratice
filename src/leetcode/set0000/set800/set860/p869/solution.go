package p869

const MAX_N = 1000000000

func reorderedPowerOf2(N int) bool {
	// 1, 2, 4, 8, 16, 32, .........
	//99999999
	PP := 1
	for i := 0; i < 30; i++ {
		if canReorder(N, PP) {
			return true
		}
		PP <<= 1
		if PP > MAX_N {
			break
		}
	}
	return false
}

func canReorder(N, M int) bool {
	cnt := make([]int, 10)
	for N > 0 {
		cnt[N%10]++
		N /= 10
	}

	for M > 0 {
		x := M % 10
		cnt[x]--
		if cnt[x] < 0 {
			return false
		}
		M /= 10
	}
	for i := 0; i < 10; i++ {
		if cnt[i] > 0 {
			return false
		}
	}
	return true
}
