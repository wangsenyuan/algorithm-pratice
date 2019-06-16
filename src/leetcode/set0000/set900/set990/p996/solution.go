package p996

import "sort"

var sqs []int64
var ii map[int64]int

func init() {
	for i := int64(0); i < 100000; i++ {
		sqs = append(sqs, i*i)
	}

	ii = make(map[int64]int)
	for i := 0; i < len(sqs); i++ {
		ii[sqs[i]] = i + 1
	}
}

func numSquarefulPerms(A []int) int {
	sort.Ints(A)

	n := len(A)
	total := 1<<uint(n) - 1
	var dfs func(num int, mask int)
	var ans int
	dfs = func(num int, mask int) {
		if mask == total {
			ans++
			return
		}
		for i := 0; i < n; i++ {
			if mask&(1<<uint(i)) > 0 {
				continue
			}

			if i > 0 && A[i] == A[i-1] && (mask&(1<<uint(i-1)) == 0) {
				continue
			}

			if num < 0 || ii[int64(num)+int64(A[i])] > 0 {
				dfs(A[i], mask|(1<<uint(i)))
			}
		}
	}

	dfs(-1, 0)

	return ans
}
