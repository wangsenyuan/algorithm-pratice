package p2085

import "testing"

func runSample(t *testing.T, street string, expect int) {
	res := minimumBuckets(street)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	str := "H..H"
	expect := 2
	runSample(t, str, expect)
}

func TestSample2(t *testing.T) {
	str := ".H.H."
	expect := 1
	runSample(t, str, expect)
}


func TestSample3(t *testing.T) {
	str := ".HHH."
	expect := -1
	runSample(t, str, expect)
}
func TestSample4(t *testing.T) {
	str := "H"
	expect := -1
	runSample(t, str, expect)
}