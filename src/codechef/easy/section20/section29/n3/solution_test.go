package main

import (
	"testing"
)

func runSample(t *testing.T, N, L, K int, S []string, expect_len int) {
	res, ops := solve(N, L, K, S)

	if len(res) != expect_len {
		t.Fatalf("Sampe expect minium length %d, but got %s", expect_len, res)
	}

	buf := make([]byte, expect_len)

	for i := 0; i < expect_len; i++ {
		buf[i] = ' '
	}

	for j, op := range ops {
		u, p := op[0], op[1]
		s := S[u]
		for i := p; i < p+L; i++ {
			if buf[i] != ' ' && buf[i] != s[i-p] {
				t.Fatalf("Sample result %v, not correct at %d", res, j)
			}
			buf[i] = s[i-p]
		}
	}

	tmp := string(buf)

	if tmp != res {
		t.Errorf("Sample result %s, not from %v", res, ops)
	}

}

func TestSample1(t *testing.T) {
	N, L, K := 3, 3, 1
	S := []string{
		"abc",
		"cde",
		"bcf",
	}
	expect_len := 3
	runSample(t, N, L, K, S, expect_len)
}

func TestSample2(t *testing.T) {
	N, L, K := 3, 3, 2
	S := []string{
		"abc",
		"cde",
		"bcf",
	}
	expect_len := 4
	runSample(t, N, L, K, S, expect_len)
}

func TestSample3(t *testing.T) {
	N, L, K := 3, 3, 3
	S := []string{
		"abc",
		"cde",
		"bcf",
	}
	expect_len := 7
	runSample(t, N, L, K, S, expect_len)
}

func TestSample4(t *testing.T) {
	N, L, K := 5, 3, 3
	S := []string{
		"abc",
		"cde",
		"efg",
		"ghi",
		"hik",
	}
	expect_len := 6
	runSample(t, N, L, K, S, expect_len)
}
