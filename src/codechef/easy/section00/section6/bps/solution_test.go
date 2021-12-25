package main

import (
	"bytes"
	"testing"
)

func TestNextPermuatation(t *testing.T) {
	p := make([]int, 4)
	for i := 0; i < 4; i++ {
		p[i] = i
	}

	res := make(map[string]bool)

	for i := 0; i < 1<<4; i++ {
		res[slice2String(p)] = true
		nextPermuatation(p)
	}
	if len(res) != 1<<4 {
		t.Errorf("expect total %d, but got %d", 1<<4, len(res))
	}
	for k := range res {
		t.Logf("res is %s", k)
	}

}

func slice2String(p []int) string {
	var buf bytes.Buffer
	for i := 0; i < len(p); i++ {
		buf.WriteByte(byte('0' + p[i]))
	}
	return buf.String()
}

func runSample(t *testing.T, n int, A []int, m int, ops [][]int, p, q int) {
	P, Q := solve(n, A, m, ops)
	if P != p || Q != q {
		t.Errorf("Sample expect %d/%d, but got %d/%d", p, q, P, Q)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{1, 2, 1}
	m := 3
	ops := [][]int{
		{1, 3},
		{1, 1},
		{3, 3},
	}
	p, q := 1, 1
	runSample(t, n, A, m, ops, p, q)
}
