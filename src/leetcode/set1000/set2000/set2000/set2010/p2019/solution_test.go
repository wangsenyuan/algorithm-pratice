package p2019

import (
	"testing"
)

func runSample(t *testing.T, s string, ans []int, expect int) {
	res := scoreOfStudents(s, ans)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "7+3*1*2"
	ans := []int{20, 13, 42}
	expect := 7
	runSample(t, s, ans, expect)
}

func TestSample2(t *testing.T) {
	s := "3+5*2"
	ans := []int{13, 0, 10, 13, 13, 16, 16}
	expect := 19
	runSample(t, s, ans, expect)
}

func TestSample3(t *testing.T) {
	s := "6+0*1"
	ans := []int{12, 9, 6, 4, 8, 6}
	expect := 10
	runSample(t, s, ans, expect)
}

func TestSample4(t *testing.T) {
	s := "1+2*3+4"
	ans := []int{13, 21, 11, 15}
	expect := 11
	runSample(t, s, ans, expect)
}
