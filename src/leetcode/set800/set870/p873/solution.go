package p873

import "sort"

func lenLongestFibSubseq(A []int) int {
	var best int
	n := len(A)

	fib := func(i, j int) int {
		cnt := 0

		for {
			k := sort.Search(n, func(k int) bool {
				return A[k] >= A[i]+A[j]
			})
			if k == n || A[k] > A[i]+A[j] {
				break
			}
			i, j = j, k
			cnt++
		}

		return cnt
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			tmp := fib(i, j)
			if tmp > 0 {
				tmp += 2
			}
			if tmp > best {
				best = tmp
			}
		}
	}

	return best
}
