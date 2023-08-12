package main

import "testing"

func runSample(t *testing.T, a, b, k int, expect bool) {
	ok, x, y := solve(a, b, k)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}

	if !ok {
		return
	}

	z := sub(x, y)

	n := len(x)

	cnt := countDigits(x)

	if cnt[0] != a || cnt[1] != b {
		t.Fatalf("Sample result\n %s\n %s\n %s\n, not correct", x[:min(10, n)], y[:min(10, n)], z[:min(10, n)])
	}

	cnt = countDigits(y)

	if cnt[0] != a || cnt[1] != b {
		t.Fatalf("Sample result\n %s\n %s\n %s\n, not correct", x[:min(10, n)], y[:min(10, n)], z[:min(10, n)])
	}

	cnt = countDigits(z)

	if cnt[1] != k {
		t.Fatalf("Sample result\n %s\n %s\n %s\n, not correct", x[:min(10, n)], y[:min(10, n)], z[:min(10, n)])
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func countDigits(s string) []int {
	cnt := make([]int, 2)
	for i := 0; i < len(s); i++ {
		cnt[int(s[i]-'0')]++
	}
	return cnt
}

func sub(a, b string) string {
	// len(a) = len(b)
	b1 := []byte(a)
	b2 := []byte(b)
	res := make([]byte, len(a))

	for i := len(a) - 1; i >= 0; i-- {
		x := int(b1[i] - '0')
		y := int(b2[i] - '0')
		if x >= y {
			res[i] = byte((x - y) + '0')
			continue
		}
		// x < y
		for j := i; j >= 0; j-- {
			if b1[j] == '1' {
				b1[j] = '0'
				break
			}
			b1[j] = '1'
		}
		res[i] = '1'
	}
	return string(res)
}

func TestSample1(t *testing.T) {
	a, b, k := 4, 2, 3
	expect := true
	runSample(t, a, b, k, expect)
}

func TestSample2(t *testing.T) {
	a, b, k := 3, 2, 1
	expect := true
	runSample(t, a, b, k, expect)
}

func TestSample3(t *testing.T) {
	a, b, k := 190006, 18, 190020
	expect := true
	runSample(t, a, b, k, expect)
}
