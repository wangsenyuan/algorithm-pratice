package p2299

func strongPasswordCheckerII(password string) bool {
	n := len(password)

	if n < 8 {
		return false
	}

	if !containsCheck(password, lower) {
		return false
	}

	if !containsCheck(password, upper) {
		return false
	}

	if !containsCheck(password, digit) {
		return false
	}

	if !containsCheck(password, special) {
		return false
	}

	for i := 1; i < n; i++ {
		if password[i] == password[i-1] {
			return false
		}
	}
	return true
}

func containsCheck(s string, fn func(byte) bool) bool {
	for i := 0; i < len(s); i++ {
		if fn(s[i]) {
			return true
		}
	}
	return false
}

func lower(x byte) bool {
	return x >= 'a' && x <= 'z'
}

func upper(x byte) bool {
	return x >= 'A' && x <= 'Z'
}

func digit(x byte) bool {
	return x >= '0' && x <= '9'
}

const SPECIAL = "!@#$%^&*()-+"

func special(x byte) bool {
	for i := 0; i < len(SPECIAL); i++ {
		if SPECIAL[i] == x {
			return true
		}
	}
	return false
}
