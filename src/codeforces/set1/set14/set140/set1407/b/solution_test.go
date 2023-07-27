package main

import "testing"

func runSample(t *testing.T, A []int, expect []int) {
	res := solve(A)

	g := make([]int, 2)

	for i := 0; i < len(A); i++ {
		g[0] = gcd(g[0], res[i])
		g[1] = gcd(g[1], expect[i])

		if g[0] != g[1] {
			t.Fatalf("Sample result %v, not correct, expect %v", res, expect)
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 5}
	expect := []int{5, 2}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 8, 2, 3}
	expect := []int{8, 2, 1, 3}
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{64, 25, 75, 100, 50}
	expect := []int{100, 50, 25, 75, 64}
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{96, 128, 88, 80, 52, 7}
	expect := []int{128, 96, 80, 88, 52, 7}
	runSample(t, A, expect)
}
func TestSample5(t *testing.T) {
	A := []int{2, 4, 8, 16, 17}
	expect := []int{17, 2, 4, 8, 16}
	runSample(t, A, expect)
}

func TestSample6(t *testing.T) {
	A := []int{757, 520, 932, 620, 590, 142, 157, 353, 218, 858}
	expect := []int{932, 520, 620, 590, 142, 218, 858, 757, 157, 353}
	runSample(t, A, expect)
}
