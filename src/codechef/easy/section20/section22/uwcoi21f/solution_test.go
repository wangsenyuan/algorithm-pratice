package main

import "testing"

func runSample(t *testing.T, n, k int) {
	res := solve(n, k)
	if len(res) != k {
		t.Fatalf("sample %d %d, result %v, not correct", n, k, res)
	}
	N := 1 << n
	F := make([]int, N)

	for i := 0; i < N; i++ {
		F[i] = -1
	}
	xx := make(map[int]int)
	yy := make(map[int]int)
	for i := 0; i < len(res); i++ {
		x, y := res[i][0], res[i][1]
		F[x] = y
		xx[x]++
		yy[y]++
	}

	if len(xx) != k {
		t.Fatalf("Sample result %v, x not distinct", res)
	}

	if len(yy) != k {
		t.Fatalf("Sample result %v, y not distinct", res)
	}

	for i := 0; i < len(res); i++ {
		y := res[i][1]
		if _, found := xx[y]; !found {
			t.Fatalf("result %v, %d, not defined in F", res, y)
		}
	}
	for x := range xx {
		if H(x) == H(F[F[x]]) {
			t.Fatalf("result %v, H(%d) == H(F(F(%d))", res, x, x)
		}
	}
}

func H(num int) int {
	return digitOneCount(num)
}

func TestSample1(t *testing.T) {
	n, k := 3, 7
	runSample(t, n, k)
}

func TestSample2(t *testing.T) {
	n, k := 3, 8
	runSample(t, n, k)
}

func TestSample3(t *testing.T) {
	n, k := 10, 500
	runSample(t, n, k)
}
