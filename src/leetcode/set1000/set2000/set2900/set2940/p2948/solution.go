package p2948

func areSimilar(mat [][]int, k int) bool {
	n := len(mat[0])
	k %= n
	for i, row := range mat {
		var tmp []int
		if i&1 == 0 {
			// 偶数行左移
			tmp = shift(row, k)
		} else {
			// 计数行右移
			tmp = shift(row, n-k)
		}
		for j := 0; j < n; j++ {
			if tmp[j] != row[j] {
				return false
			}
		}
	}
	return true
}

func shift(row []int, k int) []int {
	res := make([]int, len(row))
	copy(res, row)

	reverse(res[:k])
	reverse(res[k:])
	reverse(res)

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
