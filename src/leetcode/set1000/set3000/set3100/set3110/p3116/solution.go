package p3116

func maximumPrimeDifference(nums []int) int {
	var primes []int
	lpf := make([]int, 101)
	for i := 2; i < 100; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for j := 0; j < len(primes); j++ {
			if i*primes[j] > 100 {
				break
			}
			lpf[i*primes[j]] = primes[j]
			if i%primes[j] == 0 {
				break
			}
		}
	}

	first := -1
	var res int

	for i, num := range nums {
		if lpf[num] == num {
			if first < 0 {
				first = i
			}
			res = i - first
		}
	}
	return res

}
