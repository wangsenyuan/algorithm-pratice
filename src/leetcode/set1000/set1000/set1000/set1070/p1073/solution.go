package p1073

func addNegabinary(arr1 []int, arr2 []int) []int {
	var res []int

	i := len(arr1) - 1
	j := len(arr2) - 1
	var carry int

	for i >= 0 || j >= 0 || carry != 0 {
		if i >= 0 {
			carry += arr1[i]
			i--
		}
		if j >= 0 {
			carry += arr2[j]
			j--
		}
		res = append(res, carry&1)
		carry = -(carry >> 1)
	}
	reverse(res)
	i = 0
	for i < len(res)-1 && res[i] == 0 {
		i++
	}
	return res[i:]
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
