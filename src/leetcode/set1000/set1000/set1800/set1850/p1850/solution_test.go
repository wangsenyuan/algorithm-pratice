package p1850

import "testing"

func runSample(t *testing.T, num string, k int, expect int) {
	res := getMinSwaps(num, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := "5489355142"
	k := 4
	expect := 2
	runSample(t, num, k, expect)
}

func TestSample2(t *testing.T) {
	num := "11112"
	k := 4
	expect := 4
	runSample(t, num, k, expect)
}

func TestSample3(t *testing.T) {
	num := "00123"
	k := 1
	expect := 1
	runSample(t, num, k, expect)
}
