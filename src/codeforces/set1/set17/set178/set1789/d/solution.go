package main

import "strings"

func main() {

}

func solve(s string, t string) (int, []int) {
	lsh := func(s string, k int) string { return s[k:] + strings.Repeat("0", k) }
	rsh := func(s string, k int) string { return strings.Repeat("0", k) + s[:len(s)-k] }
	xor := func(s, t string) string {
		x := []byte(s)
		for i := range x {
			x[i] ^= t[i] ^ '0'
		}
		return string(x)
	}

	i := strings.IndexByte(s, '1')
	j := strings.IndexByte(t, '1')
	if i < 0 && j < 0 {
		return 0, nil
	}
	if i < 0 || j < 0 {
		return -1, nil
	}

	var ans []int

	if i > j {
		ans = append(ans, i-j)
		s = xor(s, lsh(s, i-j))
		i = j
	}

	k := i + 1
	for s[i+1:] != t[i+1:] {
		for s[k] == t[k] {
			k++
		}
		ans = append(ans, i-k)
		s = xor(s, rsh(s, k-i))
	}

	r := strings.LastIndexByte(s, '1')
	k = i
	for s != t {
		for s[k] == t[k] {
			k--
		}
		ans = append(ans, r-k)
		s = xor(s, lsh(s, r-k))
	}

	return len(ans), ans
}
