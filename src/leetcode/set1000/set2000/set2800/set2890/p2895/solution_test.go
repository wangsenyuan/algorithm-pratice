package p2895

import "testing"

func runSample(t *testing.T, s1, s2 string, x int, expect int) {
	res := minOperations(s1, s2, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s1 := "1100011000"
	s2 := "0101001010"
	x := 2
	expect := 4
	runSample(t, s1, s2, x, expect)
}

func TestSample2(t *testing.T) {
	s1 := "11"
	s2 := "00"
	x := 2
	expect := 1
	runSample(t, s1, s2, x, expect)
}

func TestSample3(t *testing.T) {
	s1 := "101"
	s2 := "000"
	x := 2
	expect := 2
	runSample(t, s1, s2, x, expect)
}

func TestSample4(t *testing.T) {
	s1 := "1001"
	s2 := "0110"
	x := 2
	expect := 2
	runSample(t, s1, s2, x, expect)
}

func TestSample5(t *testing.T) {
	s1 := "10110"
	s2 := "00011"
	x := 4
	expect := -1
	runSample(t, s1, s2, x, expect)
}

func TestSample6(t *testing.T) {
	s1 := "101101"
	s2 := "000000"
	x := 6
	expect := 4
	runSample(t, s1, s2, x, expect)
}

func TestSample7(t *testing.T) {
	s1 := "1100011000"
	s2 := "0101001010"
	x := 2
	expect := 4
	runSample(t, s1, s2, x, expect)
}
