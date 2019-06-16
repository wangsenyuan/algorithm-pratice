package p969

func pancakeSort(A []int) []int {
	n := len(A)

	res := make([]int, 0, 100)

	for num := n; num > 0; num-- {
		var pos = -1
		for i := 0; i < n; i++ {
			if A[i] == num {
				pos = i + 1
				break
			}
		}
		if pos != num {
			if pos > 1 {
				res = append(res, pos)
				swap(A, 0, pos)
			}
			res = append(res, num)
			swap(A, 0, num)
		}
	}

	return res
}

func swap(num []int, left, right int) {
	for i, j := left, right-1; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}
}
