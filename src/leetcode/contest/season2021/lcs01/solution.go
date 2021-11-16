package lcs01

func leastMinutes(n int) int {
	if n == 1 {
		return 1
	}
	var l int
	for 1<<l < n {
		l++
	}
	return l + 1
}
