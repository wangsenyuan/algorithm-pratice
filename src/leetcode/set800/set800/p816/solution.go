package p816

import "fmt"

func ambiguousCoordinates(S string) []string {
	S = S[1 : len(S)-1]

	n := len(S)

	res := make([]string, 0, 100)
	for i := 1; i <= n-1; i++ {
		as, bs := dot(S[:i]), dot(S[i:])

		for _, a := range as {
			for _, b := range bs {
				res = append(res, fmt.Sprintf("(%s, %s)", a, b))
			}
		}
	}

	return res
}

func dot(s string) []string {
	n := len(s)

	if n == 1 {
		return []string{s}
	}

	res := make([]string, 0, 10)
	if allowPrefix(s) {
		res = append(res, s)
	}
	for i := 1; i <= n-1; i++ {
		if allowPrefix(s[:i]) && allowSuffix(s[i:]) {
			r := make([]byte, n+1)
			copy(r[:i], s[:i])
			r[i] = '.'
			copy(r[i+1:], s[i:])
			res = append(res, string(r))
		}
	}

	return res
}

func allowPrefix(s string) bool {
	if s[0] != '0' {
		return true
	}

	return len(s) == 1
}

func allowSuffix(s string) bool {
	return s[len(s)-1] != '0'
}
