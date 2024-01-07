package p2998

import "testing"

func runSample(t *testing.T, start, finish int, limit int, s string, expect int) {
	res := numberOfPowerfulInt(int64(start), int64(finish), limit, s)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	start := 1
	finish := 6000
	limit := 4
	s := "124"
	expect := 5
	runSample(t, start, finish, limit, s, expect)
}

func TestSample2(t *testing.T) {
	start := 1000
	finish := 2000
	limit := 4
	s := "3000"
	expect := 0
	runSample(t, start, finish, limit, s, expect)
}

func TestSample3(t *testing.T) {
	start := 20
	finish := 1159
	limit := 5
	s := "20"
	expect := 8
	runSample(t, start, finish, limit, s, expect)
}
