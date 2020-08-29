package main


import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
		res := solve(n, A)
		if res != expect {
			t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
		}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 1 ,2 ,2 ,1}
	expect := 6
	runSample(t, n, A, expect)
}
func TestSample2(t *testing.T) {
	n := 6
	A := []int{2, 3, 4, 5, 6, 7}
	expect := 1
	runSample(t, n, A, expect)
}
func TestSample3(t *testing.T) {
	n := 6
	A := []int{1, 5 ,4 ,5 ,1, 9}
	expect := 6
	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 5, 6}
	expect := 4
	runSample(t, n, A, expect)
}