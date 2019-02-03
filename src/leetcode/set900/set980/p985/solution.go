package p985

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	n := len(A)
	k := len(queries)
	ans := make([]int, k)

	even := make([]bool, n)
	var sum int
	for i := 0; i < n; i++ {
		if A[i]&1 == 0 {
			even[i] = true
			sum += A[i]
		}
	}

	for i := 0; i < k; i++ {
		v := queries[i][0]
		j := queries[i][1]

		if even[j] {
			sum -= A[j]
		}
		A[j] += v
		even[j] = A[j]&1 == 0
		if even[j] {
			sum += A[j]
		}
		ans[i] = sum
	}

	return ans
}
