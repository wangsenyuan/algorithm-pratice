package p2708

func removeTrailingZeros(num string) string {
	n := len(num)

	for n > 0 && num[n-1] == '0' {
		n--
	}

	if n > 0 {
		return num[:n]
	}
	return "0"
}
