package p1980

func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	buf := make([]byte, n)

	for i := 0; i < n; i++ {
		if nums[i][i] == '0' {
			buf[i] = '1'
		} else {
			buf[i] = '0'
		}
	}
	return string(buf)
}
