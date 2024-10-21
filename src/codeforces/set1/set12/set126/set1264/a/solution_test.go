package main

import "testing"

func runSample(t *testing.T, p []int, expect []int) {
	res := solve(p)

	if sum(res) != sum(expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	if sum(expect) == 0 {
		return
	}

	g, s, b := res[0], res[1], res[2]

	if g <= 0 || s <= 0 || b <= 0 {
		t.Fatalf("Sample result %v, expect > 0", res)
	}

	if g >= s || g >= b {
		t.Fatalf("Sample result %v, expect g < s && g < b, but wrong", res)
	}

	if p[g-1] == p[g] || p[g+s-1] == p[g+s] || p[g+s+b-1] == p[g+s+b] {
		t.Fatalf("Sample result %v, expect clear bounder, but wrong", res)
	}
}

func sum(arr []int) int {
	var res int
	for _, x := range arr {
		res += x
	}

	return res
}

func TestSample1(t *testing.T) {
	p := []int{5, 4, 4, 3, 2, 2, 1, 1, 1, 1, 1, 1}
	expect := []int{1, 2, 3}
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []int{4, 3, 2, 1}
	expect := []int{0, 0, 0}
	runSample(t, p, expect)
}

func TestSample3(t *testing.T) {
	p := []int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expect := []int{2, 5, 3}
	runSample(t, p, expect)
}

func TestSample4(t *testing.T) {
	p := []int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expect := []int{2, 5, 3}
	runSample(t, p, expect)
}

func TestSample5(t *testing.T) {
	p := []int{64, 64, 63, 58, 58, 58, 58, 58, 37, 37, 37, 37, 34, 34, 28, 28, 28, 28, 28, 28, 24, 24, 19, 17, 17, 17, 17, 16, 16, 16, 16, 11}
	expect := []int{2, 6, 6}
	runSample(t, p, expect)
}

func TestSample6(t *testing.T) {
	p := []int{45, 36, 18, 18, 18, 18, 17, 17, 17, 6, 6, 6, 6, 2, 2}
	expect := []int{0, 0, 0}
	runSample(t, p, expect)
}

func TestSample7(t *testing.T) {
	p := []int{52, 52, 52, 52, 52, 52, 52, 52, 36, 36, 34, 34, 34, 33, 33, 30, 16, 16}
	expect := []int{0, 0, 0}
	runSample(t, p, expect)
}

func TestSample8(t *testing.T) {
	p := []int{28, 28, 19, 19, 19, 19, 19, 19, 19, 19, 19}
	expect := []int{0, 0, 0}
	runSample(t, p, expect)
}

func TestSample9(t *testing.T) {
	p := []int{6, 3, 3, 3, 3}
	expect := []int{0, 0, 0}
	runSample(t, p, expect)
}

func TestSample10(t *testing.T) {
	p := []int{20, 20, 20, 20, 20, 13, 13, 13, 7}
	expect := []int{0, 0, 0}
	runSample(t, p, expect)
}

func TestSample11(t *testing.T) {
	p := []int{22, 20, 12, 12, 12, 12, 12, 12}
	expect := []int{0, 0, 0}
	runSample(t, p, expect)
}

func TestSample12(t *testing.T) {
	p := []int{39, 29, 26, 24, 24, 24, 24, 8, 8, 2, 2, 2, 2}
	expect := []int{0, 0, 0}
	runSample(t, p, expect)
}

func TestSample13(t *testing.T) {
	p := []int{20, 20, 20, 20, 20, 20, 20, 20, 20}
	expect := []int{0, 0, 0}
	runSample(t, p, expect)
}

func TestSample14(t *testing.T) {
	p := []int{6, 6, 6, 6, 6, 6, 3, 3, 3, 3}
	expect := []int{0, 0, 0}
	runSample(t, p, expect)
}

func TestSample15(t *testing.T) {
	p := []int{55, 52, 52, 52, 52, 50, 49, 49, 49, 49, 47, 45, 24, 24, 22, 22, 22, 22, 22}
	expect := []int{0, 0, 0}
	runSample(t, p, expect)
}

func TestSample16(t *testing.T) {
	p := []int{35, 35, 35, 31, 31, 31, 20, 20, 20, 20, 4, 4, 4}
	expect := []int{0, 0, 0}
	runSample(t, p, expect)
}
