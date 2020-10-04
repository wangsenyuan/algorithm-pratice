package p1525

func minNumberOperations(target []int) int {
	n := len(target)

	var res int

	for i := 0; i < n; i++ {
		if i == 0 {
			res += target[i]
			continue
		}
		if target[i] > target[i-1] {
			res += target[i] - target[i-1]
		}
	}
	return res
}
