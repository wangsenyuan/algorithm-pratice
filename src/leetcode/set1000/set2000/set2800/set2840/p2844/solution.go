package p2844

func minimumOperations(num string) int {
	ans := findZeroEnding(num)
	ans = min(ans, findFiveEnding(num))
	return ans
}

func findFiveEnding(num string) int {
	var ans int
	n := len(num)
	for n > 0 && num[n-1] != '5' {
		ans++
		n--
	}
	if n <= 1 {
		return len(num)
	}
	n--

	for n > 0 && num[n-1] != '2' && num[n-1] != '7' {
		n--
		ans++
	}
	if n == 0 {
		return len(num)
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func findZeroEnding(num string) int {
	var ans int
	n := len(num)
	for n > 0 && num[n-1] != '0' {
		ans++
		n--
	}
	if n == 0 {
		return ans
	}
	if n == 1 {
		return ans + 1
	}
	n--
	for n > 0 && (num[n-1] != '0' && num[n-1] != '5') {
		ans++
		n--
	}

	return ans
}
