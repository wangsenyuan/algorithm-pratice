package p1005

import "sort"

func largestSumAfterKNegations(A []int, K int) int {
	if len(A) == 0 {
		return 0
	}
	sort.Ints(A)

	var i int

	for K > 0 && i < len(A) && A[i] < 0 {
		A[i] = -A[i]
		K--
		i++
	}

	if K%2 == 0 {
		return sum(A)
	}

	// change all element into positive nums
	sort.Ints(A)

	return sum(A) - 2*A[0]
}

func sum(arr []int) int {
	var res int
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}
