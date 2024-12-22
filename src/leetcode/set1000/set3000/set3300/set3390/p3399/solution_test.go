package p3399

import "testing"

func runSample(t *testing.T, s string, num int, expect int) {
	res := minLength(s, num)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "000001"
	num := 1
	expect := 2
	runSample(t, s, num, expect)
}

func TestSample2(t *testing.T) {
	s := "0000"
	num := 2
	expect := 1
	runSample(t, s, num, expect)
}

func TestSample3(t *testing.T) {
	s := "0101"
	num := 0
	expect := 1
	runSample(t, s, num, expect)
}

func TestSample4(t *testing.T) {
	s := "0"
	num := 1
	expect := 1
	runSample(t, s, num, expect)
}

func TestSample5(t *testing.T) {
	s := "0110"
	num := 1
	expect := 2
	runSample(t, s, num, expect)
}

func TestSample6(t *testing.T) {
	s := "00"
	num := 1
	expect := 1
	runSample(t, s, num, expect)
}

func TestSample7(t *testing.T) {
	s := "100"
	num := 1
	expect := 1
	runSample(t, s, num, expect)
}

func TestSample8(t *testing.T) {
	s := "0110"
	num := 2
	expect := 1
	runSample(t, s, num, expect)
}

func TestSample9(t *testing.T) {
	s := "011110"
	num := 2
	expect := 2
	runSample(t, s, num, expect)
}
