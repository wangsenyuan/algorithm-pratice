package p1577

import "testing"

func runSample(t *testing.T, num1, num2 []int, expect int) {
	res := numTriplets(num1, num2)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{7, 4}
	b := []int{5, 2, 8, 9}
	expect := 1
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1}
	b := []int{1, 1, 1}
	expect := 9
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{7, 7, 8, 3}
	b := []int{1, 2, 9, 7}
	expect := 2
	runSample(t, a, b, expect)
}
func TestSample4(t *testing.T) {
	a := []int{4, 7, 9, 11, 23}
	b := []int{3, 5, 1024, 12, 18}
	expect := 0
	runSample(t, a, b, expect)
}
