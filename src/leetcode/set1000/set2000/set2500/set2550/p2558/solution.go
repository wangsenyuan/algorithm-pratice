package p2558

func separateDigits(nums []int) []int {
	var res []int

	for _, num := range nums {
		var tmp []int
		for num > 0 {
			tmp = append(tmp, num%10)
			num /= 10
		}
		reverse(tmp)
		res = append(res, tmp...)
	}

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
