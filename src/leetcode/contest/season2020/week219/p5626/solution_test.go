package p5626

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := minPartitions(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "32"
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "82734"
	expect := 8
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "27346209830709182346"
	expect := 9
	runSample(t, s, expect)
}
