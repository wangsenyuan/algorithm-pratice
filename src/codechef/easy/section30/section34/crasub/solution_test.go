package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "101"
	expect := "101"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "1010"
	expect := "11"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "0000"
	expect := "0000"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "0001"
	expect := "01"
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "0010"
	expect := "01"
	runSample(t, s, expect)
}
func TestSample5(t *testing.T) {
	ss := []string{
		"0010",
		"0011",
		"00001110000",
		"00101010111",
		"10101010",
		"11110000",
		"111111",
		"10000010000",
		"100001000",
		"1000000",
		"000001",
		"00001",
		"1000",
		"10001",
		"0001001000100",
		"101",
		"1010",
		"0000",
		"0001",
	}
	expect := []string{
		"01",
		"01",
		"11100",
		"11111",
		"1111",
		"11110000",
		"111111",
		"11000",
		"110",
		"1000000",
		"01",
		"001",
		"1000",
		"101",
		"11100",
		"101",
		"11",
		"0000",
		"01",
	}
	for i := 0; i < len(ss); i++ {
		runSample(t, ss[i], expect[i])
	}
}
