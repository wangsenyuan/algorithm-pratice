package p2121

func canBeValid(s string, locked string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	// n <= 10 ^^ 5
	// at any time need open >= close, and at end open == close
	var open, close int
	for i := 0; i < n; i++ {
		if locked[i] == '1' {
			if s[i] == '(' {
				open++
			} else {
				close++
			}
		} else {
			open++
		}
		if open < close {
			return false
		}
	}
	open = 0
	close = 0
	for i := n - 1; i >= 0; i-- {
		if locked[i] == '1' {
			if s[i] == ')' {
				open++
			} else {
				close++
			}
		} else {
			open++
		}
		if open < close {
			return false
		}
	}
	return true
}
