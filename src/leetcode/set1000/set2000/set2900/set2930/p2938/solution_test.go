package p2938

import "testing"

func runSample(t *testing.T, a, b int64, n int, expect int) {
	res := maximumXorProduct(a, b, n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var a, b int64 = 12, 5
	n := 4
	expect := 98
	runSample(t, a, b, n, expect)
}

func TestSample2(t *testing.T) {
	var a, b int64 = 6, 7
	n := 5
	expect := 930
	runSample(t, a, b, n, expect)
}

func TestSample3(t *testing.T) {
	var a, b int64 = 1, 6
	n := 3
	expect := 12
	runSample(t, a, b, n, expect)
}

func TestSample4(t *testing.T) {
	var a, b int64 = 570713374625622, 553376364476768
	n := 35
	expect := 4832893
	runSample(t, a, b, n, expect)
}
