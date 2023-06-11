package p2731

import "sort"

func isFascinating(n int) bool {
	n2 := n + n
	n3 := n + n2
	var arr []int
	for n > 0 {
		arr = append(arr, n%10)
		n /= 10
	}
	for n2 > 0 {
		arr = append(arr, n2%10)
		n2 /= 10
	}
	for n3 > 0 {
		arr = append(arr, n3%10)
		n3 /= 10
	}
	if len(arr) != 9 {
		return false
	}
	sort.Ints(arr)
	if arr[0] == 0 {
		return false
	}

	for i := 0; i < 9; i++ {
		if arr[i] != i+1 {
			return false
		}
	}
	return true
}
