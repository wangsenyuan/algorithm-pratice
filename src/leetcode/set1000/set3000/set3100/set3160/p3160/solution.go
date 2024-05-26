package p3160

func occurrencesOfElement(nums []int, queries []int, x int) []int {
	var arr []int

	for i, num := range nums {
		if num == x {
			arr = append(arr, i)
		}
	}

	ans := make([]int, len(queries))

	for i, p := range queries {
		if p > len(arr) {
			ans[i] = -1
		} else {
			ans[i] = arr[p-1]
		}
	}

	return ans
}
