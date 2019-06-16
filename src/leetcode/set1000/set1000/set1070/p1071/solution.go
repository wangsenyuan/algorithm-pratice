package p1071

func gcdOfStrings(str1 string, str2 string) string {
	return gcd(str2, str1)
}

func gcd(a, b string) string {
	if a == b {
		return a
	}

	if b == "" {
		return a
	}

	if len(a) < len(b) {
		a, b = b, a
	}

	c := strMod(a, b)
	if c == "" {
		return ""
	}
	return gcd(b, c)
}

func strMod(a, b string) string {
	n := len(b)
	if a[:n] == b {
		return a[n:]
	}
	return ""
}
