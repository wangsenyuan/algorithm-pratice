package p2122

import (
	"fmt"
	"math"
)

const MOD = 100000

func abbreviateProduct(left int, right int) string {
	// left <= right
	// right <= 1000000
	var sum float64
	for i := left; i <= right; i++ {
		sum += math.Log10(float64(i))
	}

	if sum < 36 {
		var prod int64 = 1
		for i := left; i <= right; i++ {
			prod *= int64(i)
		}
		var C int
		for prod%10 == 0 {
			C++
			prod /= 10
		}
		s := fmt.Sprintf("%d", prod)
		if len(s) <= 10 {
			return fmt.Sprintf("%se%d", s, C)
		}
		f := s[:5]
		l := s[len(s)-5:]
		return output(f, l, C)
	}

	var a, b int
	var suf int64 = 1
	for i := left; i <= right; i++ {
		tmp := i
		for tmp%2 == 0 {
			a++
			tmp /= 2
		}
		for tmp%5 == 0 {
			b++
			tmp /= 5
		}
		suf = (suf * int64(tmp)) % MOD
	}

	C := min(a, b)

	for i := 0; i < a-C; i++ {
		suf = (suf * 2) % MOD
	}

	for i := 0; i < b-C; i++ {
		suf = (suf * 5) % MOD
	}

	tsuf := fmt.Sprintf("%d", suf)

	for len(tsuf) < 5 {
		tsuf = "0" + tsuf
	}

	sum -= float64(C)
	for sum > 10 {
		sum -= 6
	}
	prev := int64(math.Pow(10, sum))

	spref := fmt.Sprintf("%d", prev)[:5]

	return output(spref, tsuf, C)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func output(f, s string, c int) string {
	return fmt.Sprintf("%s...%se%d", f, s, c)
}
