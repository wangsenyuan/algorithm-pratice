package p2572

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := minOperations(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	// 54 = 32 + 16 + 4 + 2
	//   110110
	n := 54
	expect := 3
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	// 54 = 32 + 16 + 4 + 2
	//   110110
	n := 39
	expect := 3
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	// 54 = 32 + 16 + 4 + 2
	//   110110
	n := 27
	expect := 3
	runSample(t, n, expect)
}
