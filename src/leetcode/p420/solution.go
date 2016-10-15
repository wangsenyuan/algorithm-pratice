package main

import "fmt"

func main() {
	//xwfmt.Println(strongPasswordChecker("aaa123"))
	fmt.Println(strongPasswordChecker("abababababababababaaa"))
}

func strongPasswordChecker(s string) int {
	if len(s) < 6 {
		return addPassword(s)
	}

	more := len(s) - 20
	rep := 0
	sub := 0
	up, low, digit := false, false, false

	lastSame := 0
	for i := 0; i <= len(s); i++ {
		if i < len(s) {
			if isLow(s[i]) {
				low = true
			} else if isUp(s[i]) {
				up = true
			} else if isDigit(s[i]) {
				digit = true
			}
			if i == 0 || s[i] == s[i-1] {
				continue
			}
		}

		length := i - lastSame

		if more > 0 && length > 2 {
			if more >= length-2 {
				more -= length - 2
				sub += length - 2
				length = 2
			} else {
				more = 0
				sub += more
				length -= more
			}
		}

		rep += length / 3
		lastSame = i
	}

	if more > 0 {
		sub += more
	}

	add := 0
	if !up {
		add++
	}

	if !low {
		add++
	}

	if !digit {
		add++
	}

	if rep >= add {
		return rep + sub
	}
	return add + sub
}

func addPassword(s string) int {
	if len(s) <= 3 {
		return 6 - len(s)
	}

	up, low, digit := false, false, false

	for i := 0; i < len(s); i++ {
		if isLow(s[i]) {
			low = true
		} else if isUp(s[i]) {
			up = true
		} else if isDigit(s[i]) {
			digit = true
		}
	}

	add := 0
	if !up {
		add++
	}

	if !low {
		add++
	}

	if !digit {
		add++
	}

	if len(s)+add >= 6 {
		return add
	}

	return 6 - len(s)
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isLow(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func isUp(b byte) bool {
	return b >= 'A' && b <= 'Z'
}
