package p374

var guess func(int) int

func guessNumber(n int) int {
	left, right := 0, n

	for left < right {
		mid := (left + right) / 2
		res := guess(mid)
		if res == 0 {
			return mid
		}
		if res < 0 {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
