package p1768

func minOperations(boxes string) []int {
	n := len(boxes)
	// (i - j) if box[j] == 1
	var sum int
	var cnt int
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = i*cnt - sum
		if boxes[i] == '1' {
			sum += i
			cnt++
		}
	}
	cnt = 0
	sum = 0
	for i := n - 1; i >= 0; i-- {
		ans[i] += sum - i*cnt
		if boxes[i] == '1' {
			sum += i
			cnt++
		}
	}
	return ans
}
