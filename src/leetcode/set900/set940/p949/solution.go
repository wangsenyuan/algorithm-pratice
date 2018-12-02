package p949

import "fmt"

func largestTimeFromDigits(A []int) string {
	cnt := make([]int, 10)
	for i := 0; i < 4; i++ {
		cnt[A[i]]++
	}
	//pick first two first
	for i := 23; i >= 0; i-- {
		a := i / 10
		b := i % 10
		if a == b && cnt[a] > 1 || (a != b && cnt[a] > 0 && cnt[b] > 0) {
			// got one
			cnt[a]--
			cnt[b]--
			for j := 59; j >= 0; j-- {
				c := j / 10
				d := j % 10
				if c == d && cnt[c] > 1 || (c != d && cnt[c] > 0 && cnt[d] > 0) {
					return fmt.Sprintf("%d%d:%d%d", a, b, c, d)
				}
			}
			cnt[a]++
			cnt[b]++
		}
	}
	return ""
}
