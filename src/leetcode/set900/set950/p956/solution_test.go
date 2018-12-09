package p956

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := tallestBillboard(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 6}
	expect := 6
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6}
	expect := 10
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2}
	expect := 0
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{52, 61, 69, 53, 64, 69, 63, 78, 66, 65, 68, 78, 65, 68, 81, 800, 800, 800, 800, 800}
	expect := 2468
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{156, 160, 153, 149, 158, 136, 172, 147, 148, 133, 147, 147, 146, 131, 171, 165, 145, 143, 148, 145}
	expect := 1500
	runSample(t, A, expect)
}

func TestSample6(t *testing.T) {
	A := []int{92, 98, 95, 96, 87, 96, 89, 95, 100, 68, 78, 78, 84, 72, 81, 91, 900, 900, 900, 900}
	expect := 2500
	runSample(t, A, expect)
}
