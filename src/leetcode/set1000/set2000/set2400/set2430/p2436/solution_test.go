package p2436

import "testing"

func runSample(t *testing.T, time string, expect int) {
	res := countTime(time)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	time := "?5:00"
	expect := 2
	runSample(t, time, expect)
}

func TestSample2(t *testing.T) {
	time := "0?:0?"
	expect := 100
	runSample(t, time, expect)
}

func TestSample3(t *testing.T) {
	time := "??:??"
	expect := 1440
	runSample(t, time, expect)
}
