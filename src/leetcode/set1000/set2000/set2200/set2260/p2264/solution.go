package p2264

func largestGoodInteger(num string) string {
	n := len(num)
	ans := -1

	for i := 0; i+3 <= n; i++ {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			if ans < 0 || num[i] > num[ans] {
				ans = i
			}
		}
	}
	if ans < 0 {
		return ""
	}
	return num[ans : ans+3]
}
