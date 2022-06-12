package p2305

import "testing"

func runSample(t *testing.T, cookies []int, k int, expect int) {
	res := distributeCookies(cookies, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cookies := []int{8, 15, 10, 20, 8}
	k := 2
	expect := 31
	runSample(t, cookies, k, expect)
}

func TestSample2(t *testing.T) {
	cookies := []int{6, 1, 3, 2, 2, 4, 1, 2}
	k := 3
	expect := 7
	runSample(t, cookies, k, expect)
}
