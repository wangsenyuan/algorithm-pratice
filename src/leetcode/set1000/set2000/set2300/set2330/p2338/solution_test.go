package p2338

import "testing"

func runSample(t *testing.T, n int, maxValue int, expect int) {
	res := idealArrays(n, maxValue)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	maxValue := 5
	expect := 10
	runSample(t, n, maxValue, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	maxValue := 3
	expect := 11
	runSample(t, n, maxValue, expect)
}


func TestSample3(t *testing.T) {
	n := 10000
	maxValue := 10000
	expect := 22940607
	runSample(t, n, maxValue, expect)
}

func TestSample4(t *testing.T) {
	n := 100
	maxValue := 100
	expect := 959337187
	runSample(t, n, maxValue, expect)
}
