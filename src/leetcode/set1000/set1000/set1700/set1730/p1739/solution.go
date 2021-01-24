package p1739

func minimumBoxes(n int) int {
	left, right := 1, n
	N := int64(n)
	for left < right {
		mid := int64((left + right) / 2)
		res := mid * (mid + 1) * (mid + 2) / 6
		if res > 0 && res < N {
			left = int(mid + 1)
		} else {
			right = int(mid)
		}
	}

	left--
	cnt := int(int64(left) * int64(left+1) * int64(left+2) / 6)
	area := int64(left) * int64(left+1) / 2

	target := n - cnt

	left, right = 0, target

	for left < right {
		mid := (left + right) / 2
		if mid*(mid+1)/2 < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return int(area) + left
}
