package p3232

import "math"

func nonSpecialCount(l int, r int) int {
	// 平方数
	res := r - l + 1
	s := int(math.Sqrt(float64(r))) + 10
	lps := make([]int, s)
	var primes []int
	for i := 2; i < s; i++ {
		if lps[i] == 0 {
			lps[i] = i
			primes = append(primes, i)
		}
		for _, x := range primes {
			if x*i >= s {
				break
			}
			lps[x*i] = x
			if i%x == 0 {
				break
			}
		}
	}

	for _, x := range primes {
		if x*x > r {
			break
		}
		if x*x >= l {
			res--
		}
	}

	return res
}
