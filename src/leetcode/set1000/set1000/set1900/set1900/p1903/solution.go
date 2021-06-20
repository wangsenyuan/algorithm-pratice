package p1903

func largestOddNumber(num string) string {
	n := len(num)

	for i := n - 1; i >= 0; i-- {
		x := int(num[i] - '0')
		if x&1 == 1 {
			return num[:i+1]
		}
	}
	return ""
}
