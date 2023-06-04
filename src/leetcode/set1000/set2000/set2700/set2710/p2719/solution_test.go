package p2719

import "testing"

func runSample(t *testing.T, num1 string, num2 string, min_sum int, max_sum int, expect int) {
	res := count(num1, num2, min_sum, max_sum)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num1 := "1"
	num2 := "12"
	min_num := 1
	max_num := 8
	expect := 11
	runSample(t, num1, num2, min_num, max_num, expect)
}

func TestSample2(t *testing.T) {
	num1 := "1"
	num2 := "5"
	min_num := 1
	max_num := 5
	expect := 5
	runSample(t, num1, num2, min_num, max_num, expect)
}

func TestSample3(t *testing.T) {
	num1 := "10001"
	num2 := "1200010101001010"
	min_num := 100
	max_num := 300
	expect := 557080973
	runSample(t, num1, num2, min_num, max_num, expect)
}
