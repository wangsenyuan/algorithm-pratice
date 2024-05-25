package main

import "testing"

func runSample(t *testing.T, a int, b int, r int, expect int) {
	res := solve(a, b, r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 6, 0, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 0, 3, 2, 1)
}
func TestSample3(t *testing.T) {
	runSample(t, 9, 6, 10, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 92, 256, 23, 164)
}

func TestSample5(t *testing.T) {
	runSample(t, 165, 839, 201, 542)
}

func TestSample6(t *testing.T) {
	runSample(t, 1, 14, 5, 5)
}

func TestSample7(t *testing.T) {
	runSample(t, 2, 7, 2, 3)
}

func TestSample8(t *testing.T) {
	runSample(t, 96549, 34359, 13851, 37102)
}

func TestSample9(t *testing.T) {
	runSample(t, 853686404475946, 283666553522252166, 127929199446003072, 27934920819538516)
}

func TestSample10(t *testing.T) {
	runSample(t, 735268590557942972, 916721749674600979, 895150420120690183, 104449824168870225)
}
