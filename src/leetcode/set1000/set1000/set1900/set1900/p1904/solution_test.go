package p1904

import "testing"

func runSample(t *testing.T, start, finish string, expect int) {
	res := numberOfRounds(start, finish)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	start := "12:01"
	finish := "12:44"
	expect := 1
	runSample(t, start, finish, expect)
}
func TestSample2(t *testing.T) {
	start := "20:00"
	finish := "06:00"
	expect := 40
	runSample(t, start, finish, expect)
}
