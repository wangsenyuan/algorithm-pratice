package p1539

func findKthPositive(arr []int, k int) int {

	var left, right = 0, len(arr)

	for left < right {
		mid := (left + right) / 2
		//位置是 mid + 1
		//值是arr[mid]
		//左边缺失的数量是arr[mid] - (mid + 1)
		if arr[mid]-(mid+1) >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right + k
}
