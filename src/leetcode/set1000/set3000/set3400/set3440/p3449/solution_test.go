package p3449

import "testing"

func runSample(t *testing.T, points []int, m int, expect int) {
	res := maxScore(points, m)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 4}
	m := 4
	expect := 4
	runSample(t, a, m, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3}
	m := 5
	expect := 2
	runSample(t, a, m, expect)
}

func TestSample3(t *testing.T) {
	a := []int{7, 8, 13, 18}
	m := 19
	expect := 42
	runSample(t, a, m, expect)
}

func TestSample4(t *testing.T) {
	a := []int{5, 25, 15}
	m := 6
	expect := 10
	runSample(t, a, m, expect)
}

func TestSample5(t *testing.T) {
	a := []int{23, 10, 7}
	m := 25
	expect := 69
	runSample(t, a, m, expect)
}

func TestSample6(t *testing.T) {
	a := []int{2, 4}
	m := 3
	expect := 4
	runSample(t, a, m, expect)
}
