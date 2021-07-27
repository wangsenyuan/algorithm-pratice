package p1946

import "testing"

func runSample(t *testing.T, s string, change []int, expect string) {
	res := maximumNumber(s, change)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "132"
	change := []int{9, 8, 5, 0, 3, 6, 4, 2, 6, 8}
	expect := "832"
	runSample(t, s, change, expect)
}

func TestSample2(t *testing.T) {
	s := "021"
	change := []int{9, 4, 3, 5, 7, 2, 1, 9, 0, 6}
	expect := "934"
	runSample(t, s, change, expect)
}
