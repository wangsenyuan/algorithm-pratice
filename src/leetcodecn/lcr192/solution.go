package lcr192

func myAtoi(str string) int {

	n := len(str)
	var ptr int

	for ptr < n && str[ptr] == ' ' {
		ptr++
	}

	if ptr == n {
		return 0
	}

	sign := 1
	if str[ptr] == '-' {
		sign = -1
		ptr++
	} else if str[ptr] == '+' {
		ptr++
	}

	var res int
	for ptr < n && isNumber(str[ptr]) {
		x := int(str[ptr] - '0')
		res = res*10 + x
		ptr++
		if sign > 0 && res >= (1<<31)-1 {
			return (1 << 31) - 1
		}
		if sign < 0 && res >= (1<<31) {
			return -(1 << 31)
		}
	}
	res *= sign

	return res
}

func isNumber(x byte) bool {
	return x >= '0' && x <= '9'
}
